package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
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
