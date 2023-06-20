package enums

type UsernameCheckReturnType uint8

const (
	RetWithNoJSON UsernameCheckReturnType = iota
	RetTaken      UsernameCheckReturnType = iota
	RetClaimed    UsernameCheckReturnType = iota
	RetAvailable  UsernameCheckReturnType = iota
)

func (r UsernameCheckReturnType) String() string {
	return [...]string{"returned with no valid JSON", "taken", "claimed", "available"}[r]
}

type ProgramMode string

const (
	Sniping  ProgramMode = "sniping"
	Checking ProgramMode = "checking"
)

func AllProgramModesStr() []string {
	return []string{string(Sniping), string(Checking)}
}
