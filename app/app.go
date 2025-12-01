package app

import (
	"log"

	"github.com/mshirdel/sandbox/models"
	"github.com/mshirdel/sandbox/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Application struct {
	DB             *gorm.DB
	NoteRepository repository.NoteRepository
}

func New() *Application {
	// Database connection (using default values for development)
	dsn := "host=localhost user=postgres password=postgres dbname=sandbox port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		log.Printf("Continuing with nil database connection for development")
		db = nil
	}

	// Auto migrate the schema
	if db != nil {
		err = db.AutoMigrate(&models.Note{})
		if err != nil {
			log.Printf("Failed to migrate database: %v", err)
		}
	}

	// Initialize repositories
	var noteRepo repository.NoteRepository
	if db != nil {
		noteRepo = repository.NewNoteRepository(db)
	}

	return &Application{
		DB:             db,
		NoteRepository: noteRepo,
	}
}
