package main 
// package reply-challenges

import (
	"fmt"
	"io/ioutil"
    "log"
    // "os"
    "path/filepath"
)
// loops through a list of files and direccotrise
func listfile(root string) []string{
    var txtlist []string
    var list string
	
    files, err := ioutil.ReadDir(root)
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range files {
        list = f.Name()
        // fmt.Println(f.Name())
        txtlist = append(txtlist, list)
    }
    return txtlist
}

// checks wether any files that is passed from listfile() is a extfile
func checktxt(file string, num int) string{
    var txt string
    // fmt.Println(file)
    extension := filepath.Ext(file)
    if (extension == ".txt") {
        fmt.Println("file :",file," is a text file"," file no. : ",num)
        txt = file
    }
    // return extension, file
    return txt
}

// func readfile() {
	// this is the array to which the list of numbers(log statements)
	// var lis[] string

// }

func main(){
	// array of all directories, folders and files in userspecified path endpoint
	var lis[] string

	// list of text files in the user-specified path endpoint is appended to this array
	var listxtfile[] string

	var root string
	var listxt string
	
    fmt.Println("Enter path : ")
    fmt.Scanln(&root)

	lis = listfile(root)

	for g,a := range lis {
		listxt = checktxt(a,g)
		listxtfile = append(listxtfile, listxt)
	}
	// fmt.Println(listxtfile)
}