package handler

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/malahngoding/filaments/config"
	"github.com/malahngoding/filaments/utils"
)

type HandhsakeRequest struct {
	Identification string `json:"identification" xml:"identification" form:"identification"`
	Provider       string `json:"provider" xml:"provider" form:"provider"`
}

// Hello api handle
func Handshake(c *fiber.Ctx) error {

	hr := new(HandhsakeRequest)
	if err := c.BodyParser(hr); err != nil {
		return err
	}
	// ID Table See inside submitTopic function
	go utils.SubmitTopic(1, strings.Join([]string{"instead", hr.Identification}, "_"))

	s := []string{hr.Identification, "_", "_", hr.Provider}
	combined := strings.Join(s, "")

	key := []byte(config.InsteadToken())
	encrypted, err := utils.Encrypt(key, combined)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(fiber.Map{
		"messages": "Hello Future",
		"payload": fiber.Map{
			"token": encrypted,
		},
		"status": "OK",
	})
}
