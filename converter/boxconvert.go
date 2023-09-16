package converter

import (
	"log"
	"ramus/converter/ramustypes"
	"regexp"
	"strings"
)

func convertBox(src []byte) (*ramustypes.Box, error) {
	log.Println(string(src))

	regId, err := regexp.Compile("BOX \\d* ;")
	if err != nil {
		return nil, err
	}

	regDig, err := regexp.Compile("[0-9]+")
	if err != nil {
		return nil, err
	}

	regDetRef, err := regexp.Compile("DETAIL REFERENCE.*;")
	if err != nil {
		return nil, err
	}

	regRef, err := regexp.Compile("A\\d+")
	if err != nil {
		return nil, err
	}

	id := int(regDig.Find(regId.Find(src))[0]) - 48

	regName, err := regexp.Compile("'.*'")
	if err != nil {
		return nil, err
	}

	bef, aft, _ := strings.Cut(string(regName.Find(src)), ramustypes.BoxPrefix)
	name := bef + aft

	reference := string(regRef.Find(regDetRef.Find(src)))

	box := ramustypes.NewBox(id, name, reference)

	return box, nil
}
