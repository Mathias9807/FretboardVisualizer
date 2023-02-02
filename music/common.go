package music

import (
	"fmt"
	"math"
)

var NotesMap = map[string]int{
	"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11,
}

// Array of the distances between notes of a major scale, to find the interval of a given scale degree,
// sum the values of the indices less than the target degree. E.g the third has an interval of 2 + 2 = 4
// semitones
var DiatonicIntervals = []int{
	2, 2, 1, 2, 2, 2, 1,
}

// Tuning of each string from High E (0) to Low E (5) in terms of semitones away from middle C
var StringTunings = [6]int{
	16, 11, 7, 2, -3, -8,
}

// Given a certain mode of a certain key, produce an ordered list of the scale degrees where 0 is C and 11 is B. The first element will be the root note
func GetScaleDegrees(key string, mode int) [7]int {
	keyNote, exists := NotesMap[key]
	if !exists {
		panic(fmt.Sprintf("keyNote %s does not exist in list of notes", keyNote))
	}
	scale := [7]int{}
	// For each degree
	for i := 0; i < 7; i++ {
		interval := 0
		// Sum all preceeding degree distances
		for j := 0; j < i; j++ {
			interval += DiatonicIntervals[(mode+j)%7]
		}
		scale[i] = (keyNote + interval) % 12
	}
	return scale
}

// Return the frets on a given string that can play a given note
func GetFretsOfNote(note int, str int, maxFret int) []int {
	minNote := StringTunings[str]
	maxNote := minNote + maxFret
	frets := []int{}
	// Loop through i from minNote to maxNote expressed as octaves relative to `note`
	for i := math.Ceil(float64(minNote-note) / 12.0); i <= math.Floor(float64(maxNote-note)/12.0); i++ {
		frets = append(frets, int(i)*12+note-minNote)
	}
	return frets
}
