package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jtieri/Nostalgia/webserver/models"
	"github.com/jtieri/Nostalgia/webserver/repo"
	"net/http"
	"strconv"
)

func GetBlog(c *gin.Context) {
	c.File("./views/blog.html")
}

func CreateBlogPost(c *gin.Context) {
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := &models.Post{
		Title: input.Title,
		Body:  input.Body,
	}

	err := repo.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func GetAllBlogPosts(c *gin.Context) {
	blogPosts, err := repo.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogPosts)
}

func GetBlogPostById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	result, err := repo.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetBlogPostByTitle(c *gin.Context) {
	title := c.Param("title")
	result, err := repo.GetPostByTitle(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
