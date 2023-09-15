package converter

import (
	"log"
	"ramus/converter/ramustypes"
	"regexp"
	"strings"
)

type InfoType struct {
	Type  string
	BoxId int
}

func NewInfoType(src string) *InfoType {
	return &InfoType{
		Type:  string(src[1]),
		BoxId: int(src[0]) - 48,
	}
}

func convertArrow(src []byte) (arrow *ramustypes.Arrow, infoTypes []*InfoType, err error) {
	log.Println(string(src))

	regName, err := regexp.Compile("'.*'")
	if err != nil {
		return
	}

	bef, aft, _ := strings.Cut(string(regName.Find(src)), ramustypes.ArrowPrefix)
	label := bef + aft
	arrow = ramustypes.NewArrow(label)

	regInfo, err := regexp.Compile("\\d[OICM]")
	if err != nil {
		return
	}

	infos := regInfo.FindAll(src, -1)
	for _, info := range infos {
		infoTypes = append(infoTypes, NewInfoType(string(info)))
	}

	return
}
