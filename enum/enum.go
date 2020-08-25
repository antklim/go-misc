package enum

type Gender int

const (
	Unknown Gender = iota
	Male
	Female
	Other
)

func (g Gender) String() string {
	name := [...]string{
		"unknown",
		"male",
		"female",
		"other"}

	if g < Unknown || g > Other {
		return ""
	}

	return name[g]
}

type User struct {
	Gender `json:"gender" yaml:"sex"`
	Age    int `json:"age,omitempty" yaml:"age,omitempty"`
}
