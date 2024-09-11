package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
    "strconv"
	"strings"
)

// Main starts the web server and begins listening on the address and port
// specified by the LISTEN_ADDR and PORT environment variables. If these
// variables are not set, the server will listen on 0.0.0.0:8080. If the server
// fails to start, the error is printed to stdout.
func main() {
	router := setupRouter()
	port := getPort()
	if err := router.Run(getListenAddr() + ":" + strconv.Itoa(port)); err != nil {
		fmt.Println(err)
	}
}


func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(headersByRequestURI())
	router.Use(gzip.Gzip(gzip.DefaultCompression))

    router.Use(static.Serve("/", static.LocalFile("../dist", true)))

	return router
}

// getListenAddr returns the address that the server should listen on. If the
// LISTEN_ADDR environment variable is not set, the function will default to
// "0.0.0.0" and print a message to stdout.
func getListenAddr() string {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		fmt.Println("LISTEN_ADDR environment variable not set, defaulting to 0.0.0.0")
		return "0.0.0.0"
	}

	return listenAddr
}


// getPort returns the port number that the server should listen on. If the
// PORT environment variable is not set, the function will default to
// 8080 and print a message to stdout. If the value of PORT is not a valid
// number, the function will print an error message to stdout and return
// 8080.
func getPort() int {
    port := os.Getenv("PORT")
    if port == "" {
        fmt.Println("PORT environment variable not set, defaulting to 8080")
        return 8080
    }

    port = strings.TrimSpace(port) // Add this line to trim whitespace characters
    portNum, err := strconv.Atoi(port)
    if err != nil {
        fmt.Println(err)
        return 8080
    }

    return portNum
}

// headersByRequestURI is a gin middleware that sets the Cache-Control header
// on each response to "max-age=86400", which is equivalent to 24 hours. This
// is useful for serving static assets that do not change frequently, such as
// images, CSS, and JavaScript.
func headersByRequestURI() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=86400")
	}
}