package handler

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct{
	LongUrl string `json:"long_url" binding:required`
	UserId string `json:"user_id" binding:required`
}

func createShortUrl(c *gin.Context){
	var creationReqeust UrlCreationRequest
	if err := c.ShouldBindJSON(&creationReqeust); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.generateShortLink(creationReqeust.LongUrl, creationReqeust.UserId)
	store.SaveUrlMapping(shorturl, creationReqeust.LongUrl, creationReqeust.UserId)

	host:= "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message": "short url created successfully",
		"short_url": host + shortUrl
	})
}

func handleShortUrlRedirect(c *gin.Context){
	shortUrl := c.Param("shortUrl")
	initialUrl = store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}