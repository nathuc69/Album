package main

import (
	"album/backend/internal/config"
	"album/backend/internal/handlers"
	"album/backend/internal/repository"
	"album/backend/internal/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("❌ error loading .env file")
	}
	db := config.InitDB()
	defer db.Close()

	photosRepository := repository.NewPhotosRepository(db)
	photosServices := service.NewPhotosService(photosRepository)

	router := handlers.Router(photosServices)
	addr := ":8080"

	////////////////////////////////////////////////////////////////////
	////////////////// ajout de toutes les photos///////////////////////
	////////////////////////////////////////////////////////////////////

	// photosDir := "./photos"

	// // --- Parcours du dossier ---
	// err := filepath.Walk(photosDir, func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Ignorer les dossiers
	// 	if info.IsDir() {
	// 		return nil
	// 	}

	// 	// Garde uniquement les images
	// 	ext := filepath.Ext(info.Name())
	// 	switch ext {
	// 	case ".jpg", ".jpeg", ".tiff", ".JPG", ".JPEG", ".TIFF":

	// 		if _, err := os.Stat(path); err != nil {
	// 			fmt.Println("file stat error:", err)
	// 		}

	// 		photo, err := utils.ExtractMeta(path)
	// 		if err != nil {
	// 			fmt.Println("ExtractMeta error:", err)
	// 			return err
	// 		}
	// 		err = photosServices.AddPhotos(path, photo)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println("Import:", path)
	// 	}

	// 	return nil
	// })

	// if err != nil {
	// 	log.Fatal("Erreur scan dossier:", err)
	// }

	// fmt.Println("📸 Import terminé !")

	////////////////////////////////////////////////////////////////////
	////////////////// lancement du serveur ////////////////////////////
	////////////////////////////////////////////////////////////////////

	log.Printf("Server start → http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("❌ error trying to run the server: ", err)
	}
}

// package main

// import (
// 	"album/backend/utils"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	_ "github.com/mattn/go-sqlite3" // met le driver de ton SGBD
// )

// func main() {
// 	// --- Ton dossier contenant les photos ---
// 	photosDir := "./photos"

// 	// --- Connexion BDD (ici sqlite en exemple) ---
// 	db, err := sql.Open("sqlite3", "./album.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// --- Parcours du dossier ---
// 	err = filepath.Walk(photosDir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		// Ignorer les dossiers
// 		if info.IsDir() {
// 			return nil
// 		}

// 		// Garde uniquement les images
// 		ext := filepath.Ext(info.Name())
// 		switch ext {
// 		case ".jpg", ".jpeg", ".png":

// 			if _, err := os.Stat(path); err != nil {
// 				fmt.Println("file stat error:", err)
// 			}

// 			photo, err := utils.ExtractMeta(path)
// 			if err != nil {
// 				fmt.Println("ExtractMeta error:", err)
// 				return err
// 			}
// 			err = PhotosService.AddPhotoToDB(db, path, photo)
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println("Import:", path)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatal("Erreur scan dossier:", err)
// 	}

// 	fmt.Println("📸 Import terminé !")
// }
