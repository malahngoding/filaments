package handler

import (
	"fmt"
	"log"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
)

type Comments struct {
	Key        string `json:"key"`
	Comment    string `json:"comment"`
	Answer     string `json:"answer"`
	IsAnswered bool   `json:"isAnswered"`
	Lang       string `json:"lang"`
}

var dbName = "instead_comments"

// Get comments handle
func GetComments(c *fiber.Ctx) error {
	queryLang := c.Query("lang")

	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, dbName)
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}
	query := base.Query{
		{"lang": queryLang, "isAnswered": true},
	}

	// variabe to store the results
	var results []*Comments

	// fetch items
	_, err = db.Fetch(&base.FetchInput{
		Q:    query,
		Dest: &results,
	})

	if err != nil {
		fmt.Println("failed to fetch items:", err)
	}
	return c.JSON(fiber.Map{
		"messages": "Getting data from Deta",
		"payload":  fiber.Map{"comments": results},
		"status":   "OK",
	})
}

type AddCommentRequest struct {
	Message string `json:"message" xml:"message" form:"message"`
	Lang    string `json:"lang" xml:"lang" form:"lang"`
}

// Get comments handle
func AddComment(c *fiber.Ctx) error {

	acr := new(AddCommentRequest)
	if err := c.BodyParser(acr); err != nil {
		log.Println(acr.Message)
		return err
	}

	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, dbName)
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}
	// insert item in the database
	_, err = db.Insert(&Comments{
		IsAnswered: false,
		Comment:    acr.Message,
		Lang:       acr.Lang,
	})
	if err != nil {
		fmt.Println("failed to fetch items:", err)
	}
	return c.JSON(fiber.Map{
		"messages": "Successfully added comment",
		"payload":  fiber.Map{"message": acr.Message},
		"status":   "OK",
	})
}
