package main 

import ( "bufio"; "fmt"; "io/ioutil" ; "log"; "os"; "path/filepath" )

// loops through a list of files and directories
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

// checks wether any files that is passed from listfile() is a textfile
func checktxt(file string) (string, bool) {
    var txt string
    var g bool
    g = false
    // fmt.Println(file)
    extension := filepath.Ext(file)
    if (extension == ".txt") {
        fmt.Println("file :",file," is a text file")
        txt = file
        // g is true if passed file is a text file
        g = true
        // returns the passed argument only if it has a ".txt" extension
        return txt , g
    }
    // return extension, file
    // g is returned false
    return txt, g
}

// reads the contents of text files from the passed argument
func readfiles(content string) []string {
    var lis[] string
    // gets current working directory, should be removed after debugging
    path, err := os.Getwd()
    if err != nil {
        log.Println(err)
    }
    fmt.Println(path) 
	// this is the array to which the list of numbers(log statements)
	// var lis[] string

    var file, error = os.OpenFile(content, os.O_RDONLY, 0644)
    if (error != nil) {
        fmt.Println("Error in reading file")
    }

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        lis = append(lis, scanner.Text())
    }
    // close currently opened file 
    defer file.Close()
    // this is a list of the data found in the file, appended to an array
    return lis
}

func main(){
	// array of all directories, folders and files in userspecified path endpoint
	var lis[] string

	// list of text files in the user-specified path endpoint is appended to this array
	var listxtfile[] string

	var root string
	var listxt string
    var f bool
	
    fmt.Println("Enter path : ")
    fmt.Scanln(&root)

    os.Chdir(root)
    // newDir, _ := os.Getwd()
    // fmt.Println(newDir)
	lis = listfile(root)

	for _,a := range lis {
		listxt , f = checktxt(a)
        if (f==true){
            listxtfile = append(listxtfile, listxt)   
        }else {
            if(f == false){
                continue
            }
        }
	}

    for _,g := range listxtfile{
        readfiles(g)
        fmt.Println(readfiles(g))
    }
}