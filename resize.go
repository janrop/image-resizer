package main

import (
	"errors"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

func inArray(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}

func dimensionsInArgs() (dimensions int, err error) {
	for _, arg := range os.Args {
		if intVal, err2 := strconv.Atoi(arg); err2 == nil {
			dimensions = intVal
			return
		}
	}
	err = errors.New("error message")
	return
}

func inArgs(val string) (ok bool) {
	ok, _ = inArray(val, os.Args)
	return
}

func main() {

	//os.Mkdir(PATH, 0666)
	dimensions, err := dimensionsInArgs()
	if err != nil {
		fmt.Println("Usage: ")
		fmt.Println("  resize 400 # To resize all pictures with a max width and height of 400")
		fmt.Println("             -v Verbose")
		return
	}

	dirs, err := ioutil.ReadDir(".")
	check(err)

	for _, dir := range dirs {
		if !dir.IsDir() || dir.Name() == ".git" {
			continue
		}

		fmt.Println("Processing directory", dir.Name())

		files, err := ioutil.ReadDir(dir.Name() + "/originals")
		check(err)

		for _, file := range files {
			if inArgs("-v") {
				fmt.Println("  Processing file", file.Name(), file.ModTime().Format("Mon 2 Jan 2006 15:04"))
			}

			if file.ModTime().Unix() > (time.Now().Unix() - 60*60*8) {
				if inArgs("-v") {
					fmt.Println("    -> new enough")
				}

				if strings.HasSuffix(file.Name(), ".JPG") || strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".JPEG") || strings.HasSuffix(file.Name(), ".jpeg") {
					resizeJpeg(dir.Name()+"/originals", file, dimensions)
				}
			}
		}

	}
}

func resizeJpeg(folder string, fileInfo os.FileInfo, dimensions int) {
	file, err := os.Open(folder + "/" + fileInfo.Name())
	check(err)
	defer file.Close()

	newFolder := strings.Replace(folder, "originals", strconv.Itoa(dimensions), 1)
	exists, err := pathExists(newFolder)

	if !exists {
		os.Mkdir(newFolder, 0766)
	}

	img, err := jpeg.Decode(file)
	check(err)

	m := resize.Resize(uint(dimensions), 0, img, resize.Lanczos3)
	out, err := os.Create(newFolder + "/" + fileInfo.Name())
	check(err)
	defer out.Close()

	jpeg.Encode(out, m, nil)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
