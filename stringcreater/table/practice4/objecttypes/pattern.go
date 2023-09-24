package objecttypes

import "strings"

const (
	tableTitle = "| **Элемент нотации IDEF0** | **Наименование преобразуемого объекта** |\n" +
		"|---------------------------|-----------------------------------------|\n"
)

type notationElement string

const (
	entry    notationElement = "Вход"
	exit     notationElement = "Выход"
	interior notationElement = "Внутренний поток"
)

type pattern string

const (
	streamCellPattern pattern = "{StreamLabel}"
)

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(patt), "{StreamLabel}", streamLabel)
}
