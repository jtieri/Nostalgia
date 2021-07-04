package controllers

import "github.com/gin-gonic/gin"

func GetAbout(c *gin.Context) {
	c.File("./views/about.html")
}
