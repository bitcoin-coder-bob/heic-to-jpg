package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./heics")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fn := file.Name()
		if !file.IsDir() && (strings.Contains(fn, ".heic") || strings.Contains(fn, ".HEIC")) {
			dotIndex := strings.Index(fn, ".")
			filePrefix := string(fn[0:dotIndex])
			in := "./heics/" + file.Name()
			out := "./jpgs/" + filePrefix + ".jpg"
			_, err := exec.Command("heif-convert", in, out).Output()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(in + " => " + out)
			}
		}
	}
}
