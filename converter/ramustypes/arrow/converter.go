package arrow

import (
	"ramus/converter/ramustypes/arrow/types"
	"regexp"
	"strconv"
	"strings"
)

type converter struct {
	source []byte
}

func newConverter(src []byte) converter {
	return converter{
		source: src,
	}
}

func (cnv converter) convertArrowId() (arrowId int, err error) {
	arrowSegReg, _ := regexp.Compile("ARROWSEG \\d+ ;")

	digReg, _ := regexp.Compile("\\d+")

	line := arrowSegReg.Find(cnv.source)

	arrowId, err = strconv.Atoi(string(digReg.Find(line)))

	return
}

func (cnv converter) convertSourceSink(lineValue string) (*types.IOPutStruct, error) {
	var branchId int

	sourceLineReg, _ := regexp.Compile(lineValue + " (BORDER|BOX|BRANCH) (\\d*[IOCM]\\d+)?.*;")

	typeReg, _ := regexp.Compile("(BORDER|BOX|BRANCH)")

	infoReg, _ := regexp.Compile("\\d*[IOCM]\\d+")

	branchReg, _ := regexp.Compile("SOURCE BRANCH")

	line := sourceLineReg.Find(cnv.source)

	tp := string(typeReg.Find(line))
	if branchReg.Match(line) {
		digReg, _ := regexp.Compile("\\d+")

		var err error
		branchId, err = strconv.Atoi(string(digReg.Find(line)))
		if err != nil {
			return nil, err
		}
	}

	info := string(infoReg.Find(line))

	return types.NewIOPutStruct(tp, info, branchId)
}

func (cnv converter) convertLabel() (label string) {
	labelLineReg, _ := regexp.Compile("LABEL '.*' ;")

	labelReg, _ := regexp.Compile("'.*'")

	labelInfoReg, _ := regexp.Compile("\\{.*}")

	line := labelLineReg.Find(cnv.source)

	labelValue := labelReg.Find(line)
	label = string(labelInfoReg.ReplaceAll(labelValue, []byte("")))
	label = strings.ReplaceAll(label, "'", "")

	return
}
