package notationelements

import (
	"strconv"
	"strings"
)

type pattern string

const (
	tableTitle = "| **Наименование диаграммы/код** | **Вход** | **Выход** | **Механизм** | **Управление** |\n" +
		"|---------------------------------------|-----------------|------------------|---------------------|-----------------------|\n"
)

const (
	processCellPattern pattern = "{ProcessName} {ProcessReference}{ProcessId}"
	streamCellPattern  pattern = "{StreamLabel} {StreamIOPutInfo}"
)

func (patt pattern) replaceProcessNameReferenceId(processName, processReference string, processId int) string {
	value := strings.ReplaceAll(string(processCellPattern), "{ProcessName}", processName)
	value = strings.ReplaceAll(value, "{ProcessReference}", processReference)
	value = strings.ReplaceAll(value, "{ProcessId}", strconv.Itoa(processId))

	return value
}

func (patt pattern) replaceStreamLabelIOPutInfo(streamLabel, streamIOPutInfo string) string {
	value := strings.ReplaceAll(string(streamCellPattern), "{StreamLabel}", streamLabel)
	value = strings.ReplaceAll(value, "{StreamIOPutInfo}", streamIOPutInfo)

	return value
}
