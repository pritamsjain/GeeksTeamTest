package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)

/**
starts simple web server on port 8080
defines request handler for /checkTest
 */
func main() {
	r := gin.Default()
	r.POST("/checkText",handleCheckTest )
	r.Run(":8080")
}
/**
handles request for /checkText
 */
func handleCheckTest(c *gin.Context) {

	sites,siteok :=c.GetPostFormArray("Site")
	if !siteok{
		log.Println("Invalid request, Problem in site array")
		c.AbortWithStatus(http.StatusBadRequest)
	}
	searchText,textok:=c.GetPostForm("SearchText")
	if !textok{
		log.Println("Invalid request, Problem in Search Text")
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Print(searchText)

	for _,site:=range sites  {
		siteContents,err := getSiteContents(site)

		if err==nil && 0 < strings.Count(string(siteContents), searchText) {
			c.JSON(http.StatusOK, gin.H{"FoundAtSite": site})
			return
		}
	}

	c.Status(http.StatusNoContent)
	return
}
/**
method takes site url as input
returns content of that site in byte array
 */
func getSiteContents(site string) ( []byte,  error) {
	res, err := http.Get(site)
	if err != nil {
		log.Println("Error in fetching contents of ", site)
		return []byte{} ,err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
