package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count:= 0
    for _, v:= range cb[file]{
        if v{
            count++
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count:=0
    if rank < 1 || rank >8{
        return count
    }
    for _,val:= range cb{
        if val[rank-1]{ //index base 0
            count++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cnt := 0
    for key, _:= range cb{
        cnt+=CountInFile(cb, key)
    }
    return cnt
}
