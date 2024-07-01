package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	source, err := os.Open("test_file1")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer source.Close()

	dest, err := os.Create("test_file2")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dest.Close()

	srcFd := int(source.Fd())
	dstFd := int(dest.Fd())

	fi, err := source.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	_, err = syscall.Sendfile(dstFd, srcFd, nil, int(fi.Size()))
	if err != nil {
		fmt.Println("Error during sendfile:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
