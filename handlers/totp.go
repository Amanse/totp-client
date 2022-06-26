package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type TotpHandler struct {
	db *redis.Client
}

type Response struct {
	Issuers map[string]Issuer `json:"issuers"`
}

type Issuer struct {
	Accounts map[string]string `json:"accounts"`
}

func NewTotphandle(rdb *redis.Client) *TotpHandler {
	return &TotpHandler{rdb}
}

func (th *TotpHandler) GetAllCodes(c *fiber.Ctx) error {
	log.Println("Get all code")
	ctx := context.Background()
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userN := claims["username"].(string)
	dbString := fmt.Sprintf("%s%s", userN, ".secrets")
	dbStringU := fmt.Sprintf("%s%s", "users.", userN)

	val1, err1 := th.db.Exists(ctx, dbStringU).Result()
	if err1 != nil {
		panic(err1)
	}

	if val1 == 0 {
		return fiber.NewError(400, "User doesn't exist")
	}

	val, err := th.db.Get(ctx, dbString).Result()

	if err != nil {
		if err == redis.Nil {
			return fiber.NewError(400, "User has no secrets yet!")
		}
	}

	//Pass the JSON string from db directly?

	var finalData Response
	err = json.Unmarshal([]byte(val), &finalData)
	if err != nil {
		panic(err)
	}

	return c.JSON(finalData)

	// return nil
}

func (th *TotpHandler) AddCode(c *fiber.Ctx) error {
	log.Println("Adding code")
	ctx := context.Background()
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userN := claims["username"].(string)
	dbString := fmt.Sprintf("%s%s", userN, ".secrets")

	// dbString := "users.anu"

	var dataToAdd = new(Response)
	if err := c.BodyParser(dataToAdd); err != nil {
		panic(err)
	}

	th.db.Set(ctx, dbString, string(c.Body()), -1)
	return c.SendStatus(200)

}
