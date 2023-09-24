package connectiontypes

import (
	"strconv"
	"strings"
)

type pattern string

const (
	tableTitle = "| **Наименование диаграммы/код** | **Наименование потока** | **Тип связи** |\n" +
		"|---------------------------------------|--------------------------------|----------------------|\n"
)

const (
	processCellPattern pattern = "{ProcessName} {ProcessReference}{ProcessId}"
	streamCellPattern  pattern = "{StreamLabel}"
)

func (patt pattern) replaceProcessNameReferenceId(processName, processReference string, processId int) string {
	value := strings.ReplaceAll(string(processCellPattern), "{ProcessName}", processName)
	value = strings.ReplaceAll(value, "{ProcessReference}", processReference)
	value = strings.ReplaceAll(value, "{ProcessId}", strconv.Itoa(processId))

	return value
}

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(streamCellPattern), "{StreamLabel}", streamLabel)
}

type connectionType string

const (
	CONTROL connectionType = "Управление"
	IOPUT   connectionType = "Выход-Вход"
)
