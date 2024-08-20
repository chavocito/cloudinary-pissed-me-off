package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	if err != nil {
		fmt.Println("Error instantiating cloudinary client")
		return
	}
}

func buildCloudinaryClient(apiKey string, apiSecret string, cloudname string) (*cloudinary.Cloudinary, error) {
	cloudinaryUrl := fmt.Sprintf("cloudinary://%s:%s@%s", apiKey, apiSecret, cloudname)
	return cloudinary.NewFromURL(cloudinaryUrl)
}

func deleteAllAssets(tag string, cld *cloudinary.Cloudinary) {
	resp, err := cld.Admin.DeleteAssetsByTag(context.Background(), admin.DeleteAssetsByTagParams{
		Tag: tag,
	})

	if err != nil {
		fmt.Println("Error deleting assets")
		return
	}

	fmt.Println(resp.Deleted)
}

func deleteDirectory() {
	// Initialize the Cloudinary client
	cld, err := cloudinary.NewFromURL("cloudinary://:4oD7U--nIASP1GeTlU86Fdyu_sE@ebrayce")
	if err != nil {
		fmt.Println("Error instantiating cloudinary client")
		return
	}

	resp, err := cld.Admin.DeleteAllAssets(context.Background(), admin.DeleteAllAssetsParams{})
	if err != nil {
		fmt.Println("Deletion failed to complete: ", err)
	}

	// Log the delivery URL
	fmt.Println("****2. Deleted a folder****\nDelivery URL:", resp.Deleted, "\n")
}
