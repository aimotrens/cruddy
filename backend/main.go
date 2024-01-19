package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/aimotrens/cruddy/app/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	compileDate   string
	cruddyVersion string

	//go:embed static
	static embed.FS
)

func main() {
	godotenv.Load()

	fmt.Printf("Cruddy %s, compiled at %s on %s/%s\n", cruddyVersion, compileDate, runtime.GOOS, runtime.GOARCH)
	fmt.Println("Starting ...")

	rootDir := getEnvOrDefault("CRUDDY_ROOT_DIR", "./data")
	portArg := getEnvOrDefault("CRUDDY_PORT", "4231")
	apiBase := getEnvOrDefault("CRUDDY_API_BASE", "/api")
	gin.SetMode(getEnvOrDefault("GIN_MODE", gin.ReleaseMode))

	if portNumber, err := strconv.Atoi(portArg); err != nil {
		log.Fatal("Invalid port number", portArg)
	} else {
		if portNumber < 0 || portNumber > 65535 {
			log.Fatal("Invalid port number", portArg)
		}
	}

	fmt.Println("Listening on port   ", portArg)
	fmt.Println("Serving files from  ", rootDir)
	fmt.Println("API base            ", apiBase)
	fmt.Println("Gin mode            ", gin.Mode())

	g := gin.Default()

	// g.Use(cors.New(cors.Config{
	// 	AllowAllOrigins: true,
	// 	AllowMethods:    []string{"GET", "POST"},
	// 	AllowHeaders:    []string{"Origin", "Content-Type"},
	// }))

	h := api.NewHandler(rootDir)

	g.NoRoute(h.NoRoute(static))
	apiRoute := g.Group(apiBase)
	{
		apiRoute.GET("/list", h.List)
		apiRoute.GET("/download", h.Download)
		apiRoute.POST("/upload", h.Upload)
		apiRoute.POST("/delete", h.Delete)
		apiRoute.POST("/mkdir", h.Mkdir)
		apiRoute.POST("/move", h.Move)
		apiRoute.POST("/copy", h.Copy)
	}

	err := g.Run(":" + portArg)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvOrDefault(name string, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}
