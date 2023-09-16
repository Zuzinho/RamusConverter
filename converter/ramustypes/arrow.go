package ramustypes

import (
	"fmt"
	"strings"
)

type Arrow struct {
	Label string
}

type InputArrow Arrow
type OutputArrow Arrow
type ControlArrow Arrow
type MechanismArrow Arrow

func NewArrow(label string) *Arrow {
	label = strings.ReplaceAll(label, "<CR>", " ")
	label = strings.ReplaceAll(label, "'", "")
	return &Arrow{
		Label: label,
	}
}

func (arrow Arrow) StringAsList() string {
	return fmt.Sprintf("**%s**", arrow.Label)
}

const ArrowPrefix = "{LWI I 0 255 255 }"
