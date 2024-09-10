package domain

type First int
type Second int
type Result int

func (f First) Add (s Second) Result {
	return Result(int(f) + int(s))
}
