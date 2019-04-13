package main

import (
	"fmt"
	"io/ioutil"
)

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

func (df DesktopFile) toString() string {
	return fmt.Sprintf("[Desktop Entry]\nType=%s\nEncoding=%s\nName=%s\nComment=%s\nExec=%s\nIcon=%s\nTerminal=%s\n",
		df.Type, df.Encoding, df.Name, df.Comment, df.Exec, df.Icon, df.Terminal)
}

func (df DesktopFile) write(filename string) error {
	return ioutil.WriteFile(filename, []byte(df.toString()), 0666)
}
