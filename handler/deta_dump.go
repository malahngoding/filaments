package handler

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
)

type LogDump struct {
	Key  string `json:"key"`
	Log  string `json:"log"`
	Date string `json:"date"`
}

// Put deta handle
func Dump(c *fiber.Ctx) error {
	type LogBody struct {
		Action string `json:"action" xml:"action" form:"action"`
	}

	lb := new(LogBody)
	if err := c.BodyParser(lb); err != nil {
		return err
	}

	d, err := deta.New(deta.WithProjectKey(config.DetaKey()))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return err
	}

	db, err := base.New(d, "logger")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return err
	}

	str := time.Now().Local().String()

	u := &LogDump{
		Key:  fmt.Sprintf("%x", sha256.Sum256([]byte(string(str))))[:45],
		Log:  lb.Action,
		Date: str,
	}
	key, err := db.Put(u)
	return c.JSON(fiber.Map{
		"messages": "Putting data to Deta",
		"payload":  fiber.Map{"data": key},
		"status":   "OK",
	})
}
