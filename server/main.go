package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
	Number 	string	`json:"number"`
	Info 	string 	`json:"info"`
}

const DevClient = "http://localhost:3000"

func main()  {
	r := gin.Default()

	contacts := []Contact{}

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{DevClient},
        AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        // AllowOriginFunc: func(origin string) bool {
        //     return origin == "https://github.com"
        // },
        // MaxAge: 12 * time.Hour,
    }))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is working",
		})
	})

	r.GET("/contacts", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	r.POST("/contacts", func(c *gin.Context) {
		contact := &Contact{}
		
		contact.ID = len(contacts) + 1
		if err := c.BindJSON(&contact); err != nil {
			return
		}

		contacts = append(contacts, *contact)
		c.IndentedJSON(http.StatusOK, contact)
	})

	r.Use(gin.Logger())

	r.Run(":4000")
}