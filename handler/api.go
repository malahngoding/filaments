package handler

import (
	"fmt"
	"log"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
)

// Hello api handle
func Hello(c *fiber.Ctx) error {
    d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
    if err != nil {
        fmt.Println("failed to init new Deta instance:", err)
    }

    db, err := base.New(d, "base_name")
    if err != nil {
		fmt.Println(err)
    }
	um := map[string]interface{}{
        "key":      "kasdlj1",
        "username": "jimmy",
        "active":   true,
        "age":      20,
        "likes":    []string{"ramen"},
    }
    val, err := db.Put(um)
    if err != nil {
    log.Fatal("failed to put item:", err)
    }
    
	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload":  fiber.Map{"db": fmt.Sprintf("Successfully put item with key:", val)},
		"status":   "OK",
	})
}
