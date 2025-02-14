package utils

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryV2() *cloudinary.Cloudinary {
	//enviroment variables
	cloud_name := os.Getenv("CLOUD_NAME")
	api_key := os.Getenv("API_KEY")
	secret_key := os.Getenv("API_SECRET")

	if cloud_name == "" || api_key == "" || secret_key == "" {
		fmt.Println("No se encontraron las variables de entorno para cloudinary")
	} else {
		cld, _ := cloudinary.NewFromParams(cloud_name, api_key, secret_key)
		return cld
	}
	return nil
}

// Función que decodifica una imagen en base64 y la guarda en un archivo

func UploadImage(image *multipart.FileHeader) string {

	cld := CloudinaryV2()

	// Configura los parámetros de la carga
	overwrite := true
	uploadParams := uploader.UploadParams{
		Folder:    "Hotel",
		Overwrite: &overwrite,
	}

	// Realiza la carga de la imagen a Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), image, uploadParams) // Change this line
	if err != nil {
		log.Fatalf("Error uploading image: %v", err)
		return "error al subir la imagen a cloudinary"
	}

	// Retorna la URL de la imagen cargada
	return uploadResult.SecureURL

}

// delete imagen from cloudinary
func DeleteImage(public_id string) {

	cld := CloudinaryV2()
	var ctx = context.Background()
	Public_id := obtenerPublicID(public_id)
	result, err := cld.Admin.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs:    []string{"Hotel/" + Public_id},
		DeliveryType: "upload",
		AssetType:    "image",
	})
	if err != nil {
		log.Fatalf("Error deleting image: %v", err)

	}
	log.Println(result)

}

func obtenerPublicID(url string) string {
	partes := strings.Split(url, "/")
	ultimaParte := partes[len(partes)-1]
	id := strings.TrimSuffix(ultimaParte, filepath.Ext(ultimaParte))
	return id
}
