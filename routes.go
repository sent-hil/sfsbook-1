package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func ResourceShowHandler(c *gin.Context) {
	slug := c.Param("slug")

	resource, err := FindResource(slug)
	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, "resource.tmpl", gin.H{
		"Query":     slug,
		"Resources": []*Resource{resource},
	})
}

func SearchHandler(c *gin.Context) {
	query := c.Query("query")

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Query":     query,
		"Resources": []interface{}{},
	})
}
