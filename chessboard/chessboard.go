package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	_, ok := cb[file]
    if (!ok){
        return 0
    }
	count := 0
	for _, v := range cb[file] {
        if (v == true){
        	count++
        }
    }
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    rank -= 1
	if (rank < 0 || rank > 7){
        return 0
    }
	count := 0
	for _, v := range cb {
        if (v[rank] == true){
        	count++
        }
    }
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
	for _, v := range cb{
        count += len(v)
    }
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb{
       	for _, w := range v{
            if w == true{
                count++
            }
        }
    }
	return count
}
