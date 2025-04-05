package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nickalie/go-webpbin"
)

func main() {
	app := fiber.New()

	app.Post("/convert-to-webp", func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ไม่สามารถรับไฟล์ได้",
			})
		}

		// เปิดไฟล์ที่อัพโหลด
		srcFile, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "ไม่สามารถเปิดไฟล์ได้",
			})
		}
		defer srcFile.Close()

		// กำหนดชื่อไฟล์ output
		outputFile := "output/image.webp"

		// ตรวจสอบและสร้าง directory ถ้ายังไม่มี
		if err := os.MkdirAll("output", 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "ไม่สามารถสร้าง directory ได้",
			})
		}

		// แปลงไฟล์เป็น WebP
		err = webpbin.NewCWebP().
			Quality(80).
			Input(srcFile).
			OutputFile(outputFile).
			Run()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("แปลงไฟล์ล้มเหลว: %v", err),
			})
		}

		// ส่งไฟล์ WebP กลับไป
		return c.SendFile(outputFile)
	})

	// เริ่ม server
	app.Listen(":3000")
}
