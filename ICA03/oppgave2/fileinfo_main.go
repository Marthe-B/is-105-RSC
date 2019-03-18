package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var file string
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "-f" {
			file = os.Args[i+1]
		}
	}

	//file := os.Args[2]

	fi, err := os.Lstat(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Filename :", fi.Name())

	path, err := filepath.Abs(file)
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

	filemodeStr := fi.Mode().String()

	if filemodeStr[1] == "r"[0] {
		fmt.Println("User has read permission")
	} else {
		fmt.Println("User does not have read permission")
	}

	if filemodeStr[2] == "w"[0] {
		fmt.Println("User has write permission")
	} else {
		fmt.Println("User does not have write permission")
	}

	if filemodeStr[3] == "x"[0] {
		fmt.Println("User has execute permission")
	} else {
		fmt.Println("User does not have execute permission")
	}

	//fmt.Println(fi.Size())
	//fmt.Println(fi.Mode())
	//fmt.Println(fi.ModTime())
	//fmt.Println(fi.IsDir())
	//fmt.Printf("Permissions: %#o", fi.Mode().Perm())

	//Information about a file <filnavn>:
	//Size: X in bytes, kibibytes, mibibytes and gibibytes
	//Is/Is not a directory
	//Is/Is not a regular file
	//Has Unix permission bits: -rwxrwxrwx
	//Is/Is not append only
	//Is/Is not a device file
	//Is/Is not a Unix character device
	//Is/Is not a Unix block device
	//Is/Is not a symbolic link
}
