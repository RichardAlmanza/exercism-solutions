package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var counter int

	for _, value := range cb[file] {
		if value {
			counter++
		}
	}

	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	var counter int

	for _, ranks := range cb {
		if ranks[rank - 1] {
			counter++
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// This definition sounds tricky, it is time for brute force.
	// I'll miss you len(cb) * len(cb["A"])
	var counter int

	for _, ranks := range cb {
		for range ranks {
			counter++
		}
	}

	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var counter int

	for _, ranks := range cb {
		for _, square := range ranks {
			if square {
				counter++
			}
		}
	}

	return counter
}
