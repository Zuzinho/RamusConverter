package types

type Type string

const (
	INPUT     Type = "I"
	OUTPUT    Type = "O"
	CONTROL   Type = "C"
	MECHANISM Type = "M"
)

func NewArrowType(tp string) (Type, error) {
	switch tp {
	case "I":
		return INPUT, nil
	case "O":
		return OUTPUT, nil
	case "C":
		return CONTROL, nil
	case "M":
		return MECHANISM, nil
	default:
		return "", NewUnknownArrowTypeError(tp)
	}
}
