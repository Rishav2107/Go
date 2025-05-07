package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string = "example.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	defer file.Close()
	os.WriteFile(fileName, []byte("Hello, Gophers!"), 0666)
	fmt.Println("sucessfully created file")

	s := []byte("Hello, Gophers!")
	fmt.Println(s)

	file2, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while creating file: ", err)
	}
	//os.Stdout.Write(data)

	buffer := make([]byte, 32)

	//read the file content into the buffer
	for {
		data.read
	}

}
