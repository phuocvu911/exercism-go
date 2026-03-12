package highscores
import "sort"
type HighScores struct{
    s []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{s:scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.s
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.s[len(s.s)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	checker:= make([]int, len(s.s))
    copy(checker, s.s)
    sort.Slice(checker, func(i,j int) bool {
        return checker[i] > checker[j]
    })
    return checker[0]
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	checker:= make([]int, len(s.s))
    copy(checker, s.s)
    sort.Slice(checker, func(i,j int) bool {
        return checker[i] > checker[j]
    })
    if len(s.s) < 3{
        return checker
    }
    return checker[:3]
}
