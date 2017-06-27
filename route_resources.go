package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sent-hil/sfsbook-1/models"
)

func ResourceShowHandler(c *gin.Context) {
	name := c.Param("name")

	c.HTML(http.StatusOK, "resource.tmpl", gin.H{
		"Query": name,
		"Resources": []*models.Resource{
			{Name: "sfsbook"},
		},
	})
}
