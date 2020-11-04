package types

type FrettedNote struct {
	FretNumber   int
	StringNumber int
}

type Chord struct {
	Name         string
	FretCount    int
	StringCount  int
	FrettedNotes []FrettedNote
}

type ImageInfo struct {
	Width      int
	Height     int
	TargetFile string
}
