package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

type Size uint

const (
	IconSm Size = 40
	IconMd Size = 80
	IconLg Size = 160

	ImgSm Size = 320
	ImgMd Size = 640
	ImgLg Size = 980

	BgSm Size = 1280
	BgMd Size = 1920
	BgLg Size = 2840
)

func SizeMap() map[string]Size {
	sizes := map[string]Size{
		"icon-sm": IconSm,
		"icon-md": IconMd,
		"icon-lg": IconLg,

		"img-sm": ImgSm,
		"img-md": ImgMd,
		"img-lg": ImgLg,

		"bg-sm": BgSm,
		"bg-md": BgMd,
		"bg-lg": BgLg,
	}

	return sizes
}

func main() {
	sizeMap := SizeMap()

	inputDir := "./website/src/images"
	outputDir := "./website/dist/images"

	if _, err := os.Stat(outputDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(outputDir, 0750); err != nil {
				log.Fatal(err)
			}
		}
	}

	pattern := inputDir + "/*"

	filenames, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	for name, size := range sizeMap {
		for _, filename := range filenames {
			img, err := bimg.Read(filename)
			if err != nil {
				log.Fatal(err)
			}

			resized, err := bimg.Resize(img, bimg.Options{
				Width:       int(size),
				Compression: 80,
				Type:        bimg.WEBP,
			})

			if err != nil {
				log.Fatal(err)
			}

			base := filepath.Base(filename)
			ext := filepath.Ext(filename)
			baseName := base[:len(base)-len(ext)]
			outputFile := fmt.Sprintf("%s/%s-%s.webp", outputDir, baseName, name)

			if err := bimg.Write(outputFile, resized); err != nil {
				log.Fatal(err)
			}
		}
	}
}
