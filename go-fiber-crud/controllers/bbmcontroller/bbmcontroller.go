package bbmcontroller

import (
	"net/http"

	"github.com/anwarudin01/go-fiber-crud/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	// Tampung semua data yang ada pada tabel di database
	var bbm []models.Bbm
	// Cari data yang ada didalam model
	models.DB.Find(&bbm)

	// Tampilkan semua data
	return c.Status(fiber.StatusOK).JSON(bbm)
}

func Show(c *fiber.Ctx) error {
	
	// Deklarasikan Id yang dari data yang ingin ditampilkan berdasarkan parameter idnya
	id := c.Params("id")
	// Deklarasikan Struct nya
	var bbm models.Bbm

	// Cari didalam tabel yang ada di database berdasarkan id / primary keynya
	// Dan jika terjadi error, tampilkan errornya
	if err := models.DB.First(&bbm, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	// kali ini tanpa menuliskan .(Status.fiber.StatusOK) karena sudah defaultnya
	return c.JSON(bbm)
}

func Create(c *fiber.Ctx) error {
	// Deklarasikan Struct nya
	var bbm models.Bbm

	// Ambil inputannya dan
	// Mengembalikan error ketika terjadi salah input
	if err := c.BodyParser(&bbm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	// Membuat data baru dan menampilkan errornya ketika tidak berhasil
	if err := models.DB.Create(&bbm).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// mengmbalikan data yang sudah dibuat
	return c.JSON(bbm)
}

func Update(c *fiber.Ctx) error {
	// Deklarasikan Id yang dari data yang ingin ditampilkan berdasarkan parameter idnya
	id := c.Params("id")
	// Deklarasikan Struct nya
	var bbm models.Bbm

	// Ambil inputannya dan
	// Mengembalikan error ketika terjadi salah input
	if err := c.BodyParser(&bbm); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update datanya dan jika data yang diinginkan tidak ada kembalikan errornya
	if models.DB.Where("id = ?", id).Updates(&bbm).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak dapat mengupdate data",
		})
	}

	// beritau jika inputan berhasil mengupdate
	return c.JSON(fiber.Map{
		"message": "Data berhasil di update",
	})
}

func Delete(c *fiber.Ctx) error {
	// Deklarasikan Id yang dari data yang ingin ditampilkan berdasarkan parameter idnya
	id := c.Params("id")
	// Deklarasikan Struct nya
	var bbm models.Bbm

	if models.DB.Delete(&bbm, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}
	
	// beritau jika inputan berhasil menghapus
	return c.JSON(fiber.Map{
		"message": "Data berhasil di dihapus",
	})
}