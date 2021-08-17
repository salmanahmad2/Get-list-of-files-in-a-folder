package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func files(directory []fs.FileInfo) {

	//file.Name() will get all the name of the files in the folder
	for _, file := range directory {
		fmt.Println(file.Name())
	}
}

func main() {

	//pass the path of the folder as an argument to the function
	directory, err := ioutil.ReadDir("C:/Users/Awan/Documents/Practice_golang")
	if err != nil {
		fmt.Println(err)
	}
	files(directory)
}

