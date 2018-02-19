package resize

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func resizeJpeg(folder string, fileInfo os.FileInfo, dimensions int) {
	file, err := os.Open(folder + "/" + fileInfo.Name())
	check(err)
	defer file.Close()

	newFolder := strings.Replace(folder, "originals", strconv.Itoa(dimensions), 1)
	exists, err := pathExists(newFolder)

	if !exists {
		os.Mkdir(newFolder, 0755)
	}

	img, err := jpeg.Decode(file)
	check(err)

	m := resize.Resize(uint(dimensions), 0, img, resize.Lanczos3)
	out, err := os.Create(newFolder + "/" + fileInfo.Name())
	check(err)
	defer out.Close()

	jpeg.Encode(out, m, nil)
}

func resizePng(folder string, fileInfo os.FileInfo, dimensions int) {
	file, err := os.Open(folder + "/" + fileInfo.Name())
	check(err)
	defer file.Close()

	newFolder := strings.Replace(folder, "originals", strconv.Itoa(dimensions), 1)
	exists, err := pathExists(newFolder)

	if !exists {
		os.Mkdir(newFolder, 0755)
	}

	img, err := png.Decode(file)
	check(err)

	m := resize.Resize(uint(dimensions), 0, img, resize.Lanczos3)
	out, err := os.Create(newFolder + "/" + fileInfo.Name())
	check(err)
	defer out.Close()

	png.Encode(out, m)
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
