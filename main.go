package main

import (
	"log"
	"os"
	"wsmail25/config"

	"wsmail25/pkg/database"
	"wsmail25/pkg/exception"
	"wsmail25/routes"

	gcjson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder:  gcjson.Marshal,
		JSONDecoder:  gcjson.Unmarshal,
		ErrorHandler: exception.ErrHandler,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(compress.New(compress.Config{
		Level: 5,
	}))
	app.Use(cors.New(config.Cors))
	app.Use(etag.New())
	app.Use(logger.New(logger.Config{Format: "${status} - ${method} ${path}\n"}))

	mongoConn, _, err := database.NewMongoDBConnection(config.URIMONGODBWSMAIL, config.DBNAME)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	} else {
		log.Println("MongoDB connection successful!")
	}
	routes.Init(mongoConn)
	routers := routes.Router(app)
	if routers != nil {
		log.Fatalf("%+v", routers)
		return
	}
	// Listen on PORT 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // default lokal
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
