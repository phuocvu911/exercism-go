package highscores
import "slices"
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
	checker:= slices.Clone(s.s)
    slices.Sort(checker)
    return checker[len(s.s)-1]
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	checker:= slices.Clone(s.s)
    slices.Sort(checker)
    slices.Reverse(checker)
    return checker[:min(3, len(s.s))]
}
