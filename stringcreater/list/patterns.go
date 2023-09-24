package list

import "strings"

type pattern string

const (
	controlPattern   pattern = "Управляющим потоком (потоком Управления) является:"
	mechanismPattern pattern = "{ProcessType} осуществляется (механизм):"
	inputPattern     pattern = "Входами {ProcessType}а являются:"
	outputPattern    pattern = "Выходом {ProcessType}а (внутренним потоком) является:"
)

const (
	processTitle  pattern = "{ProcessType} «{ProcessName}»:"
	processInside pattern = "{ProcessType} «{ProcessName}» предполагает выполнение следующих подпроцессов:"
)

const (
	insertedProcess pattern = "«{ProcessName}»"
	insertedStream  pattern = "'{StreamLabel}'"
)

func (patt pattern) replaceProcessType(processType string) string {
	return strings.ReplaceAll(string(patt), "{ProcessType}", processType)
}

func (patt pattern) replaceProcessName(processName string) string {
	return strings.ReplaceAll(string(patt), "{ProcessName}", processName)
}

func (patt pattern) replaceProcessTypeName(processType, processName string) string {
	return strings.ReplaceAll(patt.replaceProcessType(processType), "{ProcessName}", processName)
}

func (patt pattern) replaceStreamLabel(streamLabel string) string {
	return strings.ReplaceAll(string(patt), "{StreamLabel}", streamLabel)
}

func getProcessPatterns(processType, processName string) (insideBuilder, controlBuilder, mechanismBuilder, inputBuilder, outputBuilder *strings.Builder) {
	insideBuilder = &strings.Builder{}
	controlBuilder = &strings.Builder{}
	mechanismBuilder = &strings.Builder{}
	inputBuilder = &strings.Builder{}
	outputBuilder = &strings.Builder{}

	insideBuilder.WriteString(processInside.replaceProcessTypeName(processType, processName))
	controlBuilder.WriteString(controlPattern.replaceProcessType(processType))
	mechanismBuilder.WriteString(mechanismPattern.replaceProcessType(processType))
	inputBuilder.WriteString(inputPattern.replaceProcessType(processType))
	outputBuilder.WriteString(outputPattern.replaceProcessType(processType))

	return
}
