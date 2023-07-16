package handlers

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	db *redis.Client
}

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type TokenSignIn struct {
	Success bool   `json:"success"`
	UserId  string `json:"userId"`
}

func NewAuthHand(db *redis.Client) *AuthHandler {
	return &AuthHandler{db}
}

func (ah *AuthHandler) SignUpRedis(c *fiber.Ctx) error {
	log.Println("SignupRequest")

	u := new(UserData)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	ctx := context.Background()

	userDbS := fmt.Sprintf("users.%s", u.Username)

	ifExists, err := ah.db.Exists(ctx, userDbS).Result()

	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"error": err})
	}

	if ifExists == 1 {
		log.Println("User already exists")
		return fiber.NewError(400, "User already exists")
	}

	passKey := c.Query("isPasskey")
	if passKey == "1" {
		token, err := ah.signUpPasskey(u.Username)
		if err != nil {
			return fiber.NewError(500, "Failed to get passkey token")
		}
		hPasswdS := time.Now().String()
		ah.db.Set(ctx, userDbS, hPasswdS, 0)

		return c.JSON(fiber.Map{"token": token})
	}

	//Hashing the password
	hPasswd := sha256.Sum256([]byte(u.Password))

	hPasswdS := fmt.Sprintf("%x", hPasswd)

	ah.db.Set(ctx, userDbS, hPasswdS, 0)

	claims := jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("pride"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func (ah *AuthHandler) signUpPasskey(userName string) (string, error) {
	log.Println("passkey request")
	postBody, _ := json.Marshal(map[string]string{
		"userId":   userName,
		"username": userName,
	})

	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", "https://v4.passwordless.dev/register/token", responseBody)
	if err != nil {
		return "", err
	}
	req.Header.Add("ApiSecret", "totp:secret:98219982f6e140bb8b8a9299aeb81bca")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	target := new(TokenResponse)
	json.NewDecoder(resp.Body).Decode(target)

	return target.Token, nil
}

func (ah *AuthHandler) LoginPasskey(c *fiber.Ctx) error {
	log.Println("login passkey")
	u := new(TokenResponse)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	postBody, _ := json.Marshal(map[string]string{
		"token": u.Token,
	})

	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", "https://v4.passwordless.dev/signin/verify", responseBody)
	if err != nil {
		return fiber.NewError(500, "Could not make req")
	}
	req.Header.Add("ApiSecret", "totp:secret:98219982f6e140bb8b8a9299aeb81bca")
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fiber.NewError(400, "failure to make request")
	}
	target := new(TokenSignIn)
	json.NewDecoder(resp.Body).Decode(target)

	if target.Success {
		claims := jwt.MapClaims{
			"username": target.UserId,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("pride"))
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}

	return fiber.NewError(400, "Not success login")

}

func (ah *AuthHandler) LoginRedis(c *fiber.Ctx) error {
	log.Println("Login Request")

	u := new(UserData)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	ctx := context.Background()

	result, err := ah.db.Get(ctx, fmt.Sprintf("users.%s", u.Username)).Result()

	if err != nil {
		if err == redis.Nil {
			return fiber.NewError(400, "User not found")
		}
	}

	hPasswd := sha256.Sum256([]byte(u.Password))
	hPasswdS := fmt.Sprintf("%x", hPasswd)

	if result == hPasswdS {
		claims := jwt.MapClaims{
			"username": u.Username,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("pride"))
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})

	} else {
		return fiber.NewError(400, "Username or password incorrect")
	}

}
