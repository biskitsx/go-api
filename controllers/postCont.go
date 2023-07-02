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
	newPost := new(db.Post)
	newPost.UserID = c.Locals("userId").(uint)
	newPost.Title = post.Title

	db.Db.Create(&newPost)
	// author
	user := new(db.User)
	db.Db.Find(&user, "id = ?", newPost.UserID)

	return c.JSON(
		fiber.Map{
			"post":   newPost,
			"author": user,
		},
	)
}

func UpdatePost(c *fiber.Ctx) error {
	post := new(db.Post)
	db.Db.Find(&post)
	return c.JSON(post)
}
func DeletePost(c *fiber.Ctx) error {
	posts := new(db.Post)
	db.Db.Find(&posts)
	return c.JSON(posts)
}
