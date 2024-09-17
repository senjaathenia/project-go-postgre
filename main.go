package main

import (
	"log"
	"project-go-postgre/domains"
	"project-go-postgre/pkg/config"
	"project-go-postgre/pkg/delivery"
	"project-go-postgre/pkg/repository"
	"project-go-postgre/pkg/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
    // Menghubungkan ke database menggunakan GORM
    db := config.ConnectDB()
    if db == nil {
        log.Fatal("Database connection failed")
    }
    log.Println("Database connected successfully!")

    // Inisialisasi echo framework
    e := echo.New()

    // Middleware untuk logging dan error handling
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Migrate model ke database
    migrate(db)

    // Inisialisasi repository dan usecase
    bookRepo := repository.NewBookRepository(db)
    bookUsecase := usecase.NewBookUsecase(bookRepo) // Pastikan ini mengembalikan BooksUsecase

    // Setup handler untuk endpoint buku
    delivery.NewBookHandler(e, bookUsecase) // Pastikan tipe bookUsecase sesuai

    // Menjalankan server di port 8080
    e.Logger.Fatal(e.Start(":8080"))
}

// migrate menjalankan auto-migration untuk model Buku
func migrate(db *gorm.DB) {
    // Pastikan semua tabel sudah terbuat dengan model yang digunakan
    err := db.AutoMigrate(&domains.Book{})
    if err != nil {
        log.Fatalf("Error in database migration: %v", err)
    }
    log.Println("Database migration completed!")
}
