package fileinfo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

fmt.Println("Filename :", fi.Name())

	filename := fi.Name()

	filedirectory := filepath.Dir(filename)
	path, err := filepath.Abs(filedirectory)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File path: ", path)

	fileSizeByte := fi.Size()
	fileSizeKib := fileSizeByte / 1024
	fileSizeMib := fileSizeKib / 1024
	FileSizeGib := fileSizeMib / 1024

	fmt.Println("Filesize in byte :", fileSizeByte)
	fmt.Println("Filesize in kibibytes :", fileSizeKib)
	fmt.Println("Filesize in mibibytes :", fileSizeMib)
	fmt.Println("Filesize in gibibytes :", FileSizeGib)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("Filetype: directory")
	case mode.IsRegular():
		// do file stuff
		fmt.Println("Filetype: file")
	}

	fmt.Println("Unix permission bits: ", fi.Mode())

	fmt.Printf("Permissions: %#o \n", fi.Mode().Perm())


