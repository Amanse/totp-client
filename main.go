package main

import (
	"os"

	"github.com/Amanse/totp-server/handlers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		panic("No env file")
	}
	redisUrl := os.Getenv("REDIS_URL")

	app.Use(cors.New())

	opt, err := redis.ParseURL(redisUrl)

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	ah := handlers.NewAuthHand(rdb)
	th := handlers.NewTotphandle(rdb)

	//AUTH ROUTES
	// Allowed routes
	//Signup user, takes username and password in JSON
	app.Post("/signup", ah.SignUpRedis)
	app.Post("/signup/passkey", ah.SignUpPasskey)
	app.Post("/login/passkey", ah.LoginPasskey)
	app.Post("/login", ah.LoginRedis)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("pride"),
	}))

	//CODE ROUTES
	//Private/protected
	//gets all the codes for username in URL
	//TODO: get username from jwt token and make it a protected route
	app.Get("/codes", th.GetAllCodes)
	app.Post("/codes", th.AddCode)

	app.Listen(":8080")
}
