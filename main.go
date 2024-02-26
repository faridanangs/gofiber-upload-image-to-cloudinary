package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dns := "user=postgres password=anangs port=5432 host=localhost dbname=blogs sslmode=disable TimeZone=Asia/Jakarta"
	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	app := fiber.New()

	app.Post("/upload-video", func(c *fiber.Ctx) error {
		user := User{}
		c.BodyParser(&user)
		fileHeader, _ := c.FormFile("video")
		file, _ := fileHeader.Open()
		ctx := context.Background()

		fmt.Print(file, "File ")

		cld, err := cloudinary.NewFromURL("cloudinari_url")
		if err != nil {
			log.Print(err)
		}
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
			ResourceType: "video",
			Eager:        "w_300,h_300,c_pad,ac_none|w_160,h_100,c_crop,g_south,ac_none",
			EagerAsync:   api.Bool(true),
		})
		if err != nil {
			log.Print(err, "upload")
		}
		user.Image = resp.SecureURL

		db.Model(&User{}).Create(&user)

		return c.Status(200).JSON(fiber.Map{"Data": user})
	})
	app.Get("/get", func(c *fiber.Ctx) error {
		var user []User
		db.Model(&User{}).Find(&user)
		return c.Status(200).JSON(fiber.Map{"Data": user})
	})
	app.Delete("/del/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		db.Model(&User{}).Unscoped().Delete(&User{}, "name = ?", name)
		return c.Status(200).JSON(fiber.Map{"Code": 200})
	})

	app.Listen(":3000")

}
