package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"./internal/imageHelper"
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

func getPath() (path string) {
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "/") {
			path = arg
			return
		}
	}
	path = "."
	return
}

func dimensionsInArgs() (dimensions int, err error) {
	for _, arg := range os.Args[1:] {
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

	path := getPath()
	dirs, err := ioutil.ReadDir(path)
	check(err)

	for _, dir := range dirs {
		if !dir.IsDir() || dir.Name() == ".git" {
			continue
		}

		if inArgs("-v") {
			fmt.Println("Processing directory", dir.Name())
		}

		files, err := ioutil.ReadDir(path + "/" + dir.Name() + "/originals")
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
					imageHelper.ResizeJpeg(path+"/"+dir.Name()+"/originals", file, dimensions)
				}

				if strings.HasSuffix(file.Name(), ".PNG") || strings.HasSuffix(file.Name(), ".png") {
					imageHelper.ResizePng(path+"/"+dir.Name()+"/originals", file, dimensions)
				}
			}
		}

	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
