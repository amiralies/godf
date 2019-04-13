package main

import "fmt"

type DesktopFile struct {
	Type     string
	Encoding string
	Name     string
	Comment  string
	Exec     string
	Icon     string
	Terminal string
}

const defaultDF = DesktopFile{
	Type:     "Application",
	Encoding: "UTF-8",
	Name:     "",
	Comment:  "",
	Exec:     "",
	Icon:     "",
	Terminal: "False",
}

func main() {
	fmt.Println("Hello, World!")
}
