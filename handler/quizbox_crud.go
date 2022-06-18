package handler

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

// Get CRUD handle
func GetCrud(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./quizbox.db")
	if err != nil {
		// Print error and exit if there was problem opening connection.
		log.Fatal(err)
	}
	defer db.Close()
	return c.JSON(fiber.Map{
		"messages": "Hello Future Get",
		"payload":  fiber.Map{"data": nil},
		"status":   "OK",
	})
}

// Post CRUD handle
func PostCrud(c *fiber.Ctx) error {
	sqlStmt := `
 		INSERT INTO tasks ( task, owner, checked)
		VALUES ('jon@calhoun.io', 'Jonathan', 3);
		`
	db, err := sql.Open("sqlite3", "./quizbox.db")
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	defer db.Close()

	return c.JSON(fiber.Map{
		"messages": "Hello Future Post",
		"payload":  fiber.Map{"data": nil},
		"status":   "OK",
	})
}

type Task struct {
	Id      int    `json:"id"`
	Task    string `json:"task"`
	Owner   string `json:"owner"`
	Checked int    `json:"checked"`
}

// Patch CRUD handle
func PatchCrud(c *fiber.Ctx) error {
	sqlStmt := `
		SELECT * FROM tasks;
	`
	db, err := sql.Open("sqlite3", "./quizbox.db")
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	row, err := db.Query(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	var result []Task

	for row.Next() {
		var each = Task{}
		var err = row.Scan(&each.Id, &each.Task, &each.Owner, &each.Checked)
		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, each)
	}

	defer db.Close()
	return c.JSON(fiber.Map{
		"messages": "Hello Future Patch",
		"payload":  &fiber.Map{"result": result},
		"status":   "OK",
	})
}

// Delete CRUD handle
func DeleteCrud(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"messages": "Hello Future Delete",
		"payload":  fiber.Map{"data": nil},
		"status":   "OK",
	})
}
