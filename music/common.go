package music

import "log"

var NotesMap = map[string]int{
	"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11,
}

var DiatonicIntervals = []int{
	0, 2, 4, 5, 7, 9, 11,
}

type Scale struct {
	key       string
	mode      int
	intervals []int
}

func NewScale(Key string, Mode int) Scale {
	s := Scale{}
	s.key = Key
	s.mode = Mode

	if Mode < 0 || Mode > 7 {
		log.Fatalf("NewScale called with Mode outside range 0-7: %d", Mode)
	}
	s.intervals = append(DiatonicIntervals, DiatonicIntervals...)[Mode : Mode+7]

	return s
}
