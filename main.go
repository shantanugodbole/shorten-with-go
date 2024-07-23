package main

import ("fmt" 
		"github.com/gin-gonic/gin"
		"go/url-shortener/handler"
)

func main () {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context){
		handler.CreateShortUrl(c)
		

	})

	r.GET("/:shortUrl", func(c *gin.Context){
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}