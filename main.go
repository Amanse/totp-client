package main

import (
	"github.com/Amanse/totp-server/handlers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	opt, err := redis.ParseURL("rediss://red-caosk8ns437i0d59co8g:xKAiNwq2Vq3aqjkC20p6nXBuVMWxLsPU@singapore-redis.render.com:6379")

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
