package handler

import (
	"fmt"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
)

type User struct {
	Key      string   `json:"key"`
	Username string   `json:"username"`
	Active   bool     `json:"active"`
	Age      int      `json:"age"`
	Likes    []string `json:"likes"`
}

// Put deta handle
func PutDeta(c *fiber.Ctx) error {
	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, "test_dulu")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}
	u := &User{
		Key:      "KEY_IS_SOMETHING_ELSE",
		Username: "jimmy",
		Active:   true,
		Age:      20,
		Likes:    []string{"ramen"},
	}
	key, err := db.Put(u)
	return c.JSON(fiber.Map{
		"messages": "Putting data to Deta",
		"payload":  fiber.Map{"data": key},
		"status":   "OK",
	})
}

// Get deta handle
func GetDeta(c *fiber.Ctx) error {
	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, "test_dulu")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}
	query := base.Query{}

	// variabe to store the results
	var results []*User

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
		"payload":  fiber.Map{"data": results},
		"status":   "OK",
	})
}

// Delete deta handle
func DestroyDeta(c *fiber.Ctx) error {
	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, "test_dulu")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}
	result := db.Delete("kasdlj1")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}

	return c.JSON(fiber.Map{
		"messages": "Deleting data from	 Deta",
		"payload": fiber.Map{"data": result},
		"status":  "OK",
	})
}
