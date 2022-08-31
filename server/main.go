package main

import (
	"fmt"
	"net/http"

	"github.com/Cosiamo/Contact-Cards/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Contact struct {
	gorm.Model
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
	Number 	string	`json:"number"`
	Info 	string 	`json:"info"`
}

const DevClient = "http://localhost:3000"

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "database/contacts.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&Contact{})
	fmt.Println("Database Migrated")
}

func main()  {
	r := gin.Default()
	initDatabase()

	contacts := []Contact{}
	db := database.DBConn

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

	// healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Server is working",
		})
	})

	// get all contacts
	r.GET("/contacts", func(c *gin.Context) {
		db.Find(&contacts)
		c.IndentedJSON(http.StatusOK, contacts)
	})

	// get contact by id
	r.GET("/contacts/:id", func(c *gin.Context) {
		id := c.Param("id")
		var contact Contact
		db.Find(&contact, id)
		c.IndentedJSON(http.StatusOK, contact)
	})

	// create a new contact
	r.POST("/contacts", func(c *gin.Context) {
		contact := &Contact{}
		
		// contact.ID = len(contacts) + 1
		if err := c.BindJSON(&contact); err != nil {
			return
		}

		contacts = append(contacts, *contact)
		db.Create(&contact)
		c.IndentedJSON(http.StatusOK, contact)
	})

	// delete contact
	r.DELETE("/contacts/:id", func(c *gin.Context) {
		id := c.Param("id")

		var contact Contact
		db.First(&contact, id)
		if contact.Name == "" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":"No contact found for id"})
			return
		}
		db.Delete(&contact)
		c.IndentedJSON(http.StatusOK, &contact)
	})

	r.Use(gin.Logger())
	r.Run(":4000")

	defer database.DBConn.Close()
}