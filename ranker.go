package main

import (
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func openfile() {
    f, err := os.Open("thermopylae.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    // scanner := bufio.NewScanner(f)

    // for scanner.Scan() {

    //     fmt.Println(scanner.Text())
    // }

    // if err := scanner.Err(); err != nil {
    //     log.Fatal(err)
    // }
}

func listfile(root string) []string{
    var txtlist []string
    var list string
    files, err := ioutil.ReadDir(".")
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

func main() {
    var lis[] string
    root := "d:/reply-challenges"
    lis = listfile(root)
    // fmt.Println(lis)
    for g,a := range lis {
        // fmt.Println(a,g)
        checktxt(a,g)
    }
}