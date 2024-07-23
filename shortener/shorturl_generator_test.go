package shortener

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

const userId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func testShortLinkGenerator(t *testing.T){
	link_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortened_link_1 := generateShortLink(link_1, userId)

	link_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortened_link_2 := generateShortLink(link_2, userId)

	link_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortened_link_3 := generateShortLink(link_3, userId)

	assert.Equal(t, shortened_link_1, "jTa4L57P")
	assert.Equal(t, shortened_link_2, "d66yfx7N")
	assert.Equal(t, shortened_link_3, "dhZTayYQ")

}