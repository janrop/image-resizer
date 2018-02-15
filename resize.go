package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
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

//const PATH = "resized"

func main() {

	//os.Mkdir(PATH, 0666)
	_, err := dimensionsInArgs()
	if err != nil {
		fmt.Println("Usage: ")
		fmt.Println("  resize 400 # To resize all pictures with a max width and height of 400")
		fmt.Println("             -v Verbose")
		return
	}

	dirs, err := ioutil.ReadDir(".")
	check(err)

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		fmt.Println("Processing directory", dir.Name())

		files, err := ioutil.ReadDir(dir.Name() + "/originals")
		check(err)

		for _, file := range files {
			if inArgs("-v") {
				fmt.Println("  Processing file", file.Name())
			}
		}

		// file, err := os.Open(file.Name())
		// check(err)
		// defer file.Close()

		// img, err := jpeg.Decode(file)
		// check(err)

		// newSize := float64(img.Bounds().Size().X) * 0.3

		// m := resize.Resize(uint(newSize), 0, img, resize.Lanczos3)

		// out, err := os.Create(PATH + "/" + file.Name())
		// check(err)
		// defer out.Close()

		// jpeg.Encode(out, m, nil)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
