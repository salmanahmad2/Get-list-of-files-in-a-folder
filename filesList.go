package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func files(directory []fs.FileInfo) {

	//file.Name() will get all the name of the files in the folder
	for _, file := range directory {
		fmt.Println(file.Name())
	}
}

func main() {
	file, err := ioutil.ReadDir(os.Args[1])

	//check for any error
	if err != nil {
		fmt.Println(err)
	}

	//pass the path of the folder as an argument to the function
	files(file)
}
