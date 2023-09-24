package childprocessdescription

import "strings"

const (
	tableTitle = "| **Название функции/операции** | **Краткое описание** | **Исполнитель** | **Вход** | **От кого** | **Выход** | **Кому** |\n" +
		"|-------------------------------|----------------------|-----------------|----------|-------------|-----------|----------|\n"
)

type pattern string

const (
	processNamePattern pattern = "{ProcessName}"
	processDescPattern pattern = "Процесс {ProcessName}"
	streamPattern      pattern = "{StreamLabel}"
)

const (
	entry = "Вход"
	exit  = "Выход"
)

func (patt pattern) replaceProcessName(processName string) string {
	return strings.ReplaceAll(string(patt), "{ProcessName}", processName)
}

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(patt), "{StreamLabel}", streamLabel)
}
