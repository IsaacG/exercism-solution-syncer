package highscores

import "slices"

type HighScores []int

func NewHighScores(scores []int) *HighScores {
	hs := HighScores(scores)
	return &hs
}

func (s *HighScores) Scores() []int {
	return []int(*s)
}

func (s *HighScores) Latest() int {
	return (*s)[len(*s)-1]
}

func (s *HighScores) PersonalBest() int {
	scores := slices.Clone(*s)
	slices.Sort(scores)
	return scores[len(scores)-1]
}

func (s *HighScores) TopThree() []int {
	scores := slices.Clone(*s)
	slices.Sort(scores)
	slices.Reverse(scores)
	return scores[:min(3, len(scores))]
}
