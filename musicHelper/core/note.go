package core

type Note int

const (
	Ab Note = 0
	A Note = 1
	Bb Note = 2
	B Note = 3
	C Note = 4
	Db Note = 5
	D Note = 6
	Eb Note = 7
	E Note = 8
	F Note = 9
	Gb Note = 10
	G Note = 11
)
var vals = [...]string {
	"Ab","A","Bb","B",
	"C","Db","D","Eb",
	"E","F","Gb","G",
} 

func (n *Note) String() string {
	if *n > G {
		return ""
	}
	return vals[*n]
}