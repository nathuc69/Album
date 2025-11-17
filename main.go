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
