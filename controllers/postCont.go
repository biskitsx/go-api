package controllers

import (
	"github.com/biskitsx/go-api/db"
	"github.com/gofiber/fiber/v2"
)

type TitleInput struct {
	Title string `json:"title"`
}

func GetPosts(c *fiber.Ctx) error {
	posts := new([]db.Post)
	db.Db.Find(&posts)
	return c.JSON(posts)
}
func GetPost(c *fiber.Ctx) error {
	post := new(db.Post)
	db.Db.Find(&post, "id = ?", c.Params("id"))
	if post.ID == 0 {
		return c.JSON(fiber.Map{"msg": "user not found"})
	}
	return c.JSON(post)
}
func CreatePost(c *fiber.Ctx) error {
	post := new(TitleInput)
	if err := c.BodyParser(post); err != nil {
		return c.JSON(fiber.Map{"err": "error from input"})
	}

	newPost := &db.Post{
		UserID: c.Locals("userId").(uint),
		Title:  post.Title,
	}

	db.Db.Create(newPost)

	// Retrieve the user with the updated posts
	user := new(db.User)
	db.Db.Preload("Posts").First(&user, newPost.UserID)

	return c.JSON(fiber.Map{
		"post":   newPost,
		"author": user,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	updatedPost := new(db.Post)
	db.Db.Find(&updatedPost, "id = ?", c.Params("id"))
	if updatedPost.ID == 0 { // not found
		return c.JSON(fiber.Map{"err": "no user with this id"})
	}

	post := new(TitleInput)
	if err := c.BodyParser(post); err != nil {
		return c.JSON(fiber.Map{"err": "error from input"})
	}

	updatedPost.Title = post.Title
	db.Db.Save(&updatedPost)
	return c.JSON(updatedPost)
}

func DeletePost(c *fiber.Ctx) error {
	post := new(db.Post)
	user := new(db.User)
	db.Db.Delete(&post, "id = ?", c.Params("id"))
	db.Db.First(&user, "id = ?", c.Locals("userId").(uint))

	var updatedPosts []db.Post
	for _, p := range user.Posts {
		if p.ID != post.ID {
			updatedPosts = append(updatedPosts, p)
		}
	}
	user.Posts = updatedPosts
	db.Db.Save(&user)
	return c.JSON(post)
}
