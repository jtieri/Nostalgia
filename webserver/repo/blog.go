package repo

import (
	"github.com/jtieri/Nostalgia/webserver/app"
	"github.com/jtieri/Nostalgia/webserver/models"
)

func CreatePost(post *models.Post) error {
	result := app.WebApp.DB.Create(post)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetAllPosts() ([]*models.Post, error) {
	var blogPosts []*models.Post

	result := app.WebApp.DB.Find(&blogPosts)
	if result.Error != nil {
		return nil, result.Error
	}

	return blogPosts, nil
}

func GetPostByID(ID int) (*models.Post, error) {
	var post *models.Post

	result := app.WebApp.DB.First(&post, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func GetPostByTitle(title string) (*models.Post, error) {
	var post *models.Post

	result := app.WebApp.DB.First(&post, "title = ?", title)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func UpdatePostTitle() {}

func UpdatePostBody() {}

func DeletePost() {}
