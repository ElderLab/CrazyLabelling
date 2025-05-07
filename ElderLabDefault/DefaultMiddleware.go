package ElderLabDefault

import (
	"os"

	"github.com/ElderLab/CrazyLabelling"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func DefaultMiddleware(app *fiber.App) {
	var Output *os.File
	if CrazyLabelling.IsProd {
		name := os.Getenv("ELDERLAB_LOG_FILE")
		if name == "" {
			name = "/var/log/elderlab/elderlab.log"
		}
		file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		Output = file
	} else {
		Output = os.Stdout
	}

	// Add request ID middleware for traceability
	app.Use(requestid.New())

	// Add request logging middleware with enhanced security information
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${latency} - ${status} - ${method} - ${ip} - ${reqHeader:user-agent} - ReqID:${locals:requestid}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Output:     Output,
	}))

	// Add CORS middleware to allow cross-origin requests
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (configure more restrictively in production)
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	if CrazyLabelling.IsProd {
		log.SetLevel(log.LevelWarn)
	} else {
		log.SetLevel(log.LevelDebug)
	}
	log.SetOutput(Output)
}
