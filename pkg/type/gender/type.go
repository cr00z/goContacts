package gender

type Gender int8

const (
	UNKNOWN Gender = 0
	MALE    Gender = 1
	FEMALE  Gender = 2
)

func New(gender int8) Gender {
	switch gender {
	case 1:
		return MALE
	case 2:
		return FEMALE
	default:
		return UNKNOWN
	}
}

func (g Gender) String() string {
	switch g {
	case MALE:
		return "Male"
	case FEMALE:
		return "Female"
	default:
		return "unknown"
	}
}
