package resources

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/dsi-hackaton/messages"
	"github.com/tonyalaribe/dsi-hackaton/models"
)

type Location struct {
	Base
}

func (p Location) Get(c *gin.Context) {
	log.Println("get")
	client := models.Client{Path: "./waste.bolt"}
	err := client.Open()
	if err != nil {
		log.Println(err)
	}
	defer client.DB.Close()

	locations, err := models.GetAll(client.DB)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, locations)
}

func (p Location) Post(c *gin.Context) {
	log.Println("location post")
	client := models.Client{Path: "./waste.bolt"}
	err := client.Open()
	if err != nil {
		log.Println(err)
	}
	defer client.DB.Close()

	location := models.Location{}
	err = c.BindJSON(&location)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
	}

	err = models.Create(client.DB, location)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
	}
	c.JSON(http.StatusOK, messages.Success)
}

func (p Location) Put(c *gin.Context) {
	log.Println("put")
	client := models.Client{Path: "./waste.bolt"}
	err := client.Open()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
	}
	defer client.DB.Close()
	location := models.Location{}
	err = c.BindJSON(&location)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
	}

	err = models.UpdateLocation(client.DB, location)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
	}
	c.JSON(http.StatusOK, messages.Success)
}
