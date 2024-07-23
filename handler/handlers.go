package handler

import(
	"go/url-shortener/shortener"
	"go/url-shortener/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type UrlCreationRequest struct{
	LongUrl string `json:"long_url" binding:required`
	UserId string `json:"user_id" binding:required`
}

func CreateShortUrl(c *gin.Context){
	var creationReqeust UrlCreationRequest
	if err := c.ShouldBindJSON(&creationReqeust); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("inside this function")
	shortUrl := shortener.GenerateShortLink(creationReqeust.LongUrl, creationReqeust.UserId)
	store.SaveUrlMapping(shortUrl, creationReqeust.LongUrl, creationReqeust.UserId)

	host:= "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message": "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context){
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}