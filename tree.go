package tonho

// Location represents a location in the source code.
type Location interface {
	Start() int
	End() int

	// Text gets the text of the file at the location.
	Text() string

	// File gets the file name of the location.
	File() string
}

type Tree interface {
	// Location gets the location of the tree.
	Location() Location
}
