package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/graphics-go/graphics"
)

func resizeImage(img image.Image, width, height int) image.Image {
    dstImage := image.NewRGBA(image.Rect(0, 0, width, height))
    graphics.Scale(dstImage, img)
    return dstImage
}

func PngToJpg(pngImagePath string, outFilePath string) (string, error) {
	pngImgFile, err := os.Open(pngImagePath)

	if err != nil {
	fmt.Println("Arquivo png não encontrado!")
	return "", err
	}

	defer pngImgFile.Close()

	// create image from PNG file
	imgSrc, err := png.Decode(pngImgFile)

	if err != nil {
	fmt.Println(err)
	return "", err
	}

	// create a new Image with the same dimension of PNG image
	imgr := resizeImage(imgSrc, 1200, 1200)
	newImg := image.NewRGBA(imgr.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.RGBA{255, 255, 255, 1.0}}, image.Point{}, draw.Src)
	offsetX := (1200 - imgSrc.Bounds().Dx()) / 12
	offsetY := (1200 - imgSrc.Bounds().Dy()) / 20
	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), imgr, image.Pt(offsetX, offsetY), draw.Over)

	// create new out JPEG file
	jpgFilePath := strings.Replace(pngImagePath, "png", "jpg", 1)
	jpgImgFile, err := os.Create(outFilePath)

	if err != nil {
	fmt.Printf("Não foi possivel criar %s !", outFilePath)
	fmt.Println(err)
	return "", err
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 80

	// convert newImage to JPEG encoded byte and save to jpgImgFile
	// with quality = 80
	err = jpeg.Encode(jpgImgFile, newImg, &opt)

	if err != nil {
	fmt.Println(err)
	return "", err
	}

	fmt.Println("Converted PNG file to JPEG file")
	return jpgFilePath, nil
	}

	func main() {
		// Read the PNG image file

		regex := regexp.MustCompile(`_clipped_rev_1\.png`)
		entryPath := "./fotos"
		outPath := "./out"

		if _, err := os.Stat(outPath); os.IsNotExist(err) {
			println("Diretorio criado")
				err :=os.Mkdir("./out", os.ModePerm)
				if err != nil {
					println("Erro ao criar diretorio")
				}
		}

		err := filepath.WalkDir(entryPath, func(path string, d fs.DirEntry, err error) error {

		nameToSave := regex.ReplaceAllString(d.Name(), ".jpg")
		fullEntryPath := entryPath +"/"+ d.Name()
		fullOutPath := outPath +"/"+ nameToSave
		println(d.Name())

		if d.Name() != "fotos"{
		jpegBytes, error := PngToJpg(fullEntryPath, fullOutPath)
			if error != nil {
			log.Fatalf("Erro ao converter imagem: %s", err)
		}
		println(jpegBytes)
	}
			return nil
		})
		if err != nil {
			log.Fatalf("Impossivel ler esse diretorio: %s", err)
		}

	}


