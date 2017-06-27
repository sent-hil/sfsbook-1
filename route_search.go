package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {
	query := c.Query("query")

	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Query":     query,
		"Resources": []interface{}{},
	})
}
