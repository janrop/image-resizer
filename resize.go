package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/janrop/image-resizer/internal/imageHelper"
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

func inArgs(val string) (ok bool) {
	ok, _ = inArray(val, os.Args[1:])
	return
}

func main() {

	defer func() {
		fmt.Println(recover())
	}()

	dimensionsInArgs, j := inArray("-w", os.Args)

	if !dimensionsInArgs {
		fmt.Println("Usage: ")
		fmt.Println("  resize  ")
		fmt.Println("          -w (required) [int] Resize all pictures to a max width of [int]")
		fmt.Println("          -t [int] Only resize files older than [int] seconds")
		fmt.Println("          -v Verbose output")
		fmt.Println("          [target_directory] to resize files in (default: ./)")
		return
	}

	dimensions, err := strconv.Atoi(os.Args[j+1])
	check(err)

	timeSpecified, i := inArray("-t", os.Args)
	var fileEditedSince int64
	if timeSpecified {
		fileEditedSince, _ = strconv.ParseInt(os.Args[i+1], 10, 64)
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

			if !timeSpecified || file.ModTime().Unix() > (time.Now().Unix()-fileEditedSince) {
				if inArgs("-v") {
					if timeSpecified {
						fmt.Println("    -> new enough, resizing")
					} else {
						fmt.Println("    -> resizing")
					}
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
