package checker

import (
	"log"
	"mime/multipart"
	"strings"
)

func IsCorrectFormat(header *multipart.FileHeader) (bool, error) {
	name := (*header).Filename
	log.Printf("checking file '%s' format", name)

	splt := strings.Split(name, ".")
	if len(splt) == 0 {
		log.Println("file name is empty")
		return false, NewIncorrectFormatErr("")
	}

	format := splt[len(splt)-1]
	return format == "idl", NewIncorrectFormatErr(format)
}
