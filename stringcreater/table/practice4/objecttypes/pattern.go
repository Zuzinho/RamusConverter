package objecttypes

import "strings"

const (
	tableTitle = "| **Элемент нотации IDEF0** | **Наименование преобразуемого объекта** |\n" +
		"|---------------------------|-----------------------------------------|\n"
)

type notationElement string

const (
	ENTRY    notationElement = "Вход"
	EXIT     notationElement = "Выход"
	INTERIOR notationElement = "Внутренний поток"
)

type pattern string

const (
	streamCellPattern pattern = "{StreamLabel}"
)

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(patt), "{StreamLabel}", streamLabel)
}
