package types

import (
	"regexp"
	"strconv"
	"strings"
)

type Reference struct {
	Letter       string
	ReferenceIds []int
}

func NewReference(src []byte) (*Reference, error) {
	letterReg, _ := regexp.Compile("[A-Z]")
	digReg, _ := regexp.Compile("\\d")

	letter := string(letterReg.Find(src))

	if letter == "" {
		return nil, NewInvalidReferenceError(string(src))
	}

	digs := digReg.FindAll(src, -1)
	if len(digs) == 0 {
		return nil, NewInvalidReferenceError(string(src))
	}

	ids := make([]int, 0)
	for _, v := range digs[:len(digs)-1] {
		id, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}

		if id == 0 {
			break
		}

		ids = append(ids, id)
	}

	return &Reference{
		Letter:       letter,
		ReferenceIds: ids,
	}, nil
}

func (ref Reference) String() string {
	builder := strings.Builder{}
	builder.WriteString(ref.Letter)

	for _, v := range ref.ReferenceIds {
		builder.WriteString(strconv.Itoa(v))
	}

	return builder.String()
}
