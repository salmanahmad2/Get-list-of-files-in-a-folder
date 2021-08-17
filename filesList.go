package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func getFiles(directory []fs.FileInfo, err error) {
	// the following code will check if there is any error in the path speicfied
	if err != nil {
		fmt.Println(err)
	}

	//file.Name() will get all the name of the files in the folder
	for _, file := range directory {
		fmt.Println(file.Name())
	}
}

func main() {
	//pass the path of the folder as an argument to the function
	getFiles(ioutil.ReadDir("C:/Users/Awan/Documents/Practice_golang"))
}
