// Package yacht scores a roll in Yacht.
package yacht

/*
Ones            1 × number of ones      1 1 1 4 5 scores 3
Twos            2 × number of twos      2 2 3 4 5 scores 4
Threes          3 × number of threes    3 3 3 3 3 scores 15
Fours           4 × number of fours     1 2 3 3 5 scores 0
Fives           5 × number of fives     5 1 5 2 5 scores 15
Sixes           6 × number of sixes     2 3 4 5 6 scores 6
Full House      Total of the dice       3 3 3 5 5 scores 19
Four of a Kind  Total of the four dice  4 4 4 4 6 scores 16
Little Straight 30 points               1 2 3 4 5 scores 30 
Big Straight    30 points               2 3 4 5 6 scores 30
Choice          Sum of the dice         2 3 3 4 6 scores 18
Yacht           50 points               4 4 4 4 4 scores 50
*/

var scorer = map[string]turn{
	"ones": singles{1},
	"twos": singles{2},
	"threes": singles{3},
	"fours": singles{4},
	"fives": singles{5},
	"sixes": singles{6},
	"full house": fullHouse{},
	"four of a kind": fourOfAKind{},
	"little straight": straight{6},
	"big straight": straight{1},
	"choice": choice{},
	"yacht": yacht{},
}

func toMap (dice []int) map[int]int {
	m := make(map[int]int)
	for _, d := range dice {
		if _, ok := m[d]; !ok {
			m[d] = 1
		} else {
			m[d]++
		}
	}
	return m
}

func sum (dice []int) (s int) {
	for _, d := range dice {
		s += d
	}
	return s
}

// turn encapsulates information needed to score a turn.
type turn interface {
	score (dice []int) int
}

type singles struct {
	num int
}
type fullHouse struct {}
type fourOfAKind struct {}
type straight struct {
	missing int
}
type choice struct {}
type yacht struct {}

func (t singles) score(dice []int) (score int) {
	for _, d := range dice {
		if d == t.num {
			score += t.num
		}
	}
	return score
}


func (t fullHouse) score(dice []int) int {
	m := toMap(dice)
	if len(m) == 2 && (m[dice[0]] == 2 || m[dice[0]] == 3) {
		return sum(dice)
	}
	return 0
}

func (t fourOfAKind) score(dice []int) int {
	m := toMap(dice)
	if m[dice[0]] >= 4 {
		return 4 * dice[0]
	} else if m[dice[1]] >= 4 {
		return 4 * dice[1]
	}
	return 0
}

func (t straight) score(dice []int) int {
	m := toMap(dice)
	if _, ok := m[t.missing]; len(m) == 5 && !ok {
		return 30
	}
	return 0
}

func (t choice) score(dice []int) int {
	return sum(dice)
}

func (t yacht) score(dice []int) int {
	if len(toMap(dice)) == 1 {
		return 50
	}
	return 0
}

// Score scores a hand of dice and a category.
func Score(dice []int, category string) int {
	return scorer[category].score(dice)
}
