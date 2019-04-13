package main

type DesktopFile struct {
	Type     string
	Encoding string
	Name     string
	Comment  string
	Exec     string
	Icon     string
	Terminal string
}

var defaultDF = DesktopFile{
	Type:     "Application",
	Encoding: "UTF-8",
	Name:     "",
	Comment:  "",
	Exec:     "",
	Icon:     "",
	Terminal: "False",
}

func newDF() DesktopFile {
	return defaultDF
}
