package config

import (
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"http://localhost:3000",
	"http://localhost:5000",
	"http://localhost:1300",
	
}

var headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"token",
	"Login",
	"email",
	"tanggal_tahun",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
}

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowHeaders:     strings.Join(headers[:], ","),
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
