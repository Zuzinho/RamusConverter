package processdescription

import "strings"

const (
	tableTitle = "| **Название подпроцесса** | **Краткое описание** | **Исполнитель** | **Вход** | **Выход** |\n" +
		"|--------------------------|----------------------|-----------------|----------|-----------|\n"
)

type pattern string

const (
	processNamePattern pattern = "{ProcessName}"
	processDescPattern pattern = "Процесс {ProcessName}"
	streamPattern      pattern = "{StreamLabel}"
)

func (patt pattern) replaceProcessName(processName string) string {
	return strings.ReplaceAll(string(patt), "{ProcessName}", processName)
}

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(patt), "{StreamLabel}", streamLabel)
}
