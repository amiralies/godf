package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func read(description string) string {
	fmt.Print(description + " ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}

	return strings.TrimSuffix(input, "\n")
}

func getValue(value, defaultValue string) string {
	if value != "" {
		return value
	}

	return defaultValue
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	plainDF := newDF()

	filename := read("Enter filname eg: qutebrowser.desktop")

	dfType := read(fmt.Sprintf("Type: [%s]", plainDF.Type))
	dfEncoding := read(fmt.Sprintf("Encoding: [%s]", plainDF.Encoding))
	dfName := read(fmt.Sprintf("Name:"))
	dfComment := read(fmt.Sprintf("Comment:"))
	dfExec := read(fmt.Sprintf("Exec:"))
	dfIcon := read(fmt.Sprintf("Icon:"))
	dfTerminal := read(fmt.Sprintf("Terminal: [%s]", plainDF.Terminal))

	df := DesktopFile{
		Type:     getValue(dfType, plainDF.Type),
		Encoding: getValue(dfEncoding, plainDF.Encoding),
		Name:     getValue(dfName, plainDF.Name),
		Comment:  getValue(dfComment, plainDF.Comment),
		Exec:     getValue(dfExec, plainDF.Exec),
		Icon:     getValue(dfIcon, plainDF.Icon),
		Terminal: getValue(dfTerminal, plainDF.Terminal),
	}

	writeErr := df.write(filename)
	if writeErr != nil {
		panic(writeErr)
	}
}
