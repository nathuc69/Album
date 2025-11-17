package main

import (
	"fmt"
	"os"

	"album/backend/utils"
)

func main() {
	path := "photos/IMG_0125.JPEG"

	// Check that the file exists (common cause of failures)
	if _, err := os.Stat(path); err != nil {
		fmt.Println("file stat error:", err)
	}

	photo, err := utils.ExtractMeta(path)
	if err != nil {
		fmt.Println("ExtractMeta error:", err)
		return
	}

	// Print the full struct and specific fields for debugging
	fmt.Printf("photo: %+v\n", photo)
	fmt.Println("CameraMake:", photo.CameraMake)
}
