package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glibs/gin-webserver"
)

func main() {
	host := "127.0.0.1:81"
	server := InitializeServer(host)
	server.Start()
	log.Println("Gin web server started on " + host)
	for {
		// Hold program open to test web server
	}
	// After the server is not useful
	// server.Stop()
}

func InitializeServer(host string) (server *network.WebServer) {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.LoadHTMLGlob("./templates/*")
	engine.Static("/public/css/", "./public/css")
	engine.Static("/public/js/", "./public/js/")
	engine.Static("/public/fonts/", "./public/fonts/")
	engine.Static("/public/img/", "./public/img/")

	engine.GET("/", getIndex)
	engine.GET("/about", getAbout)
	engine.GET("/contact", getContact)

	return network.InitializeWebServer(engine, host)
}

func getIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"variableName": "value",
	})
}

func getAbout(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"variableName": "value",
	})
}

func getContact(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"variableName": "value",
	})
}
