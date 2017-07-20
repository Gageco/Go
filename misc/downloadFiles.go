package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
  //"strconv"
)

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func main() {
	url := "http://images.gotinder.com/588a6d1f0a702f71402e6d10/07eb323e-c882-484f-80c5-18a05098668d.jpg"
	downloadFromUrl(url)
	os.Rename("./07eb323e-c882-484f-80c5-18a05098668d.jpg", ".01.jpg")
}
