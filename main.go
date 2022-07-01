package main

// Never ever panic in the routing the request

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/dwivedi-ritik/filehost-go/db"
	"github.com/dwivedi-ritik/filehost-go/models"
	"github.com/gofiber/fiber/v2"
)

func getFileHashValue(arg string) string {
	h := sha256.New()
	h.Write([]byte(arg))
	s := hex.EncodeToString(h.Sum(nil))
	return s
}

func Encode(num int) string {
	ALPHABET := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base := len(ALPHABET)
	var s strings.Builder
	for num > 0 {
		c := string(ALPHABET[num%base])
		s.WriteString(c)
		num = num / base
	}

	org := s.String()
	var newEncoded strings.Builder
	for i := len(org) - 1; i >= 0; i-- {
		newEncoded.WriteString(string(org[i]))
	}
	return newEncoded.String()

}

func Decode(encoded string) int {
	ALPHABET := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base := len(ALPHABET)
	n := 0
	for i := 0; i < len(encoded); i++ {
		n = n*base + strings.Index(ALPHABET, string(encoded[i]))
	}
	return n
}

func home(c *fiber.Ctx) error {
	allEntry, err := db.GetAllRow()
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(allEntry)
}

func getFile(c *fiber.Ctx) error {
	params := c.Params("file")
	id_ := Decode(params)
	var GetFile models.File
	GetFile.ID = uint(id_)
	fetchedFile, err := db.GetRow(GetFile)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Download("./Uploads/"+fetchedFile.FileHash, fetchedFile.FileName)

}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("data")

	if err != nil {
		return c.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	fileHash := getFileHashValue(file.Filename)

	newFileEntry := models.File{FileName: file.Filename, FileHash: fileHash, FileSize: file.Size, DownloadCount: 0}

	id, _ := db.AddRow(&newFileEntry)

	c.SaveFile(file, fmt.Sprintf("./Uploads/%s", fileHash))

	new_url := Encode(int(id))
	return c.SendString(c.Hostname() + "/" + new_url)
}

func main() {
	db.MakeMigration()
	app := fiber.New()

	app.Get("/all", home)

	app.Post("/", uploadFile)
	app.Get("/:file", getFile)
	log.Fatal(app.Listen(":8000"))

}
