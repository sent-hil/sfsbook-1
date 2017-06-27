package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sent-hil/sfsbook-1/models"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func ResourceShowHandler(c *gin.Context) {
	name := c.Param("name")

	c.HTML(http.StatusOK, "resource.tmpl", gin.H{
		"Query": name,
		"Resources": []*models.Resource{
			{Name: "sfsbook"},
		},
	})
}

func SearchHandler(c *gin.Context) {
	query := c.Query("query")

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Query":     query,
		"Resources": []interface{}{},
	})
}
