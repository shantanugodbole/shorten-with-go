package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init(){
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T){
	assert.True(t, testStoreService.redisClient != nil)
}

func testInsertionAndRetrieval(t *testing.T){
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	SaveUrlMapping(shortUrl, initialLink, userUUId)

	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)

}