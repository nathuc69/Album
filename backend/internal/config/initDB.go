package config

import (
	repositories "album/backend/internal/repository"
	"database/sql"
	"fmt"

	//"forum/internal/repositories"
	"os"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	//connexion à la BdD:
	dbPath := os.Getenv("FORUM_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("❌ error opening database:", err)
	}
	//vérification de la connexion:
	if err = db.Ping(); err != nil {
		log.Fatal("❌ error connecting to database:", err)
	}
	//exécution des migrations (ie modifications structurelles de la BdD):
	if err := repositories.RunMigrations(db, "migration"); err != nil {
		log.Fatal("❌ error migrating:", err)
		fmt.Println("test erreur")
	}
	fmt.Println("test")
	return db
}
