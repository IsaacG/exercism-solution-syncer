// zebra solution based on kahgoh's solution in Elixir.
// Each "identifier" is represented by one bit allowing a uint16 to represent all the detail of a house.
// Five houses (or a street) can be represented by a [5]uint16.
package zebra

import rand "math/rand/v2"

const (
	ShiftNationality = 0
	ShiftPet         = 3
	ShiftColor       = 6
	ShiftDrink       = 9
	ShiftSmoke       = 12
	MaskNationality  = 0b111 << ShiftNationality
	MaskPet          = 0b111 << ShiftPet
	MaskColor        = 0b111 << ShiftColor
	MaskDrink        = 0b111 << ShiftDrink
	MaskSmoke        = 0b111 << ShiftSmoke
	MaskAll          = MaskNationality | MaskPet | MaskColor | MaskDrink | MaskSmoke
	Englishman       = 1 << ShiftNationality
	Spaniard         = 2 << ShiftNationality
	Japanese         = 3 << ShiftNationality
	Ukrainian        = 4 << ShiftNationality
	Norwegian        = 5 << ShiftNationality
	Dog              = 1 << ShiftPet
	Fox              = 2 << ShiftPet
	Snails           = 3 << ShiftPet
	Horse            = 4 << ShiftPet
	Zebra            = 5 << ShiftPet
	Red              = 1 << ShiftColor
	Ivory            = 2 << ShiftColor
	Green            = 3 << ShiftColor
	Blue             = 4 << ShiftColor
	Yellow           = 5 << ShiftColor
	Tea              = 1 << ShiftDrink
	Coffee           = 2 << ShiftDrink
	OrangeJuice      = 3 << ShiftDrink
	Milk             = 4 << ShiftDrink
	Water            = 5 << ShiftDrink
	OldGold          = 1 << ShiftSmoke
	Kools            = 2 << ShiftSmoke
	LuckyStrike      = 3 << ShiftSmoke
	Parliaments      = 4 << ShiftSmoke
	Chesterfields    = 5 << ShiftSmoke
)

var nations = map[uint16]string{
	Englishman: "Englishman",
	Spaniard:   "Spaniard",
	Japanese:   "Japanese",
	Ukrainian:  "Ukrainian",
	Norwegian:  "Norwegian",
}

// Street has 5 houses, each a uint16.
type Street [5]uint16

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// Return the string representation of the nationality of a house which meets a criteria.
func (s Street) nationality(mask, value uint16) string {
	house := s[s.index(mask, value)]
	for k, v := range nations {
		if house&MaskNationality == k {
			return v
		}
	}
	return ""
}

// index returns the index of a house on a street which has a given property.
func (s Street) index(mask, value uint16) int {
	for idx, house := range s {
		if house&mask == value {
			return idx
		}
	}
	return -1
}

// swap returns a copy of a street which has a property swapped between two houses.
func swap(street Street, idxA, idxB int, mask uint16) Street {
	revMask := MaskAll ^ mask
	copy := street
	copy[idxA] = street[idxA]&revMask | street[idxB]&mask
	copy[idxB] = street[idxB]&revMask | street[idxA]&mask
	return copy
}

// houseRule checks if a single house has two properties.
// If a single house meets the requirements, return ok.
// Otherwise return new streets where that rule holds true.
func houseRule(street Street, maskA, valA, maskB, valB uint16) ([]Street, bool) {
	a := street.index(maskA, valA)
	b := street.index(maskB, valB)
	if a == b {
		return nil, true
	}
	return []Street{swap(street, a, b, maskA), swap(street, a, b, maskB)}, false
}

// neighborRule checks if two adjacent houses have two properties.
// If the street meets the requirements, return ok.
// Otherwise return new streets where that rule holds true.
func neighborRule(street Street, maskA, valA, maskB, valB uint16) ([]Street, bool) {
	a := street.index(maskA, valA)
	b := street.index(maskB, valB)
	if abs(a-b) == 1 {
		return nil, true
	}
	alt := make([]Street, 0, 4)
	if a > 0 {
		alt = append(alt, swap(street, a-1, b, maskB))
	}
	if a < 4 {
		alt = append(alt, swap(street, a+1, b, maskB))
	}
	if b > 0 {
		alt = append(alt, swap(street, b-1, a, maskA))
	}
	if b < 4 {
		alt = append(alt, swap(street, b+1, a, maskA))
	}
	return alt, false
}

func ruleGreenIvory(street Street) ([]Street, bool) {
	a := street.index(MaskColor, Green)
	b := street.index(MaskColor, Ivory)
	if a == b+1 {
		return nil, true
	}
	alt := []Street{}
	if a > 0 {
		alt = append(alt, swap(street, a-1, b, MaskColor))
	}
	if b < 4 {
		alt = append(alt, swap(street, b+1, a, MaskColor))
	}
	return alt, false
}

func ruleMilkMiddle(street Street) ([]Street, bool) {
	a := 2
	b := street.index(MaskDrink, Milk)
	if a == b {
		return nil, true
	}
	return []Street{swap(street, a, b, MaskDrink)}, false
}

func ruleNorwayFirst(street Street) ([]Street, bool) {
	a := 0
	b := street.index(MaskNationality, Norwegian)
	if a == b {
		return nil, true
	}
	return []Street{swap(street, a, b, MaskNationality)}, false
}

var rules = []func(Street) ([]Street, bool){
	func(s Street) ([]Street, bool) { return houseRule(s, MaskNationality, Englishman, MaskColor, Red) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskNationality, Spaniard, MaskPet, Dog) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskDrink, Coffee, MaskColor, Green) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskNationality, Ukrainian, MaskDrink, Tea) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskSmoke, OldGold, MaskPet, Snails) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskSmoke, Kools, MaskColor, Yellow) },
	func(s Street) ([]Street, bool) { return houseRule(s, MaskSmoke, LuckyStrike, MaskDrink, OrangeJuice) },
	func(s Street) ([]Street, bool) {
		return houseRule(s, MaskNationality, Japanese, MaskSmoke, Parliaments)
	},
	ruleGreenIvory,
	ruleMilkMiddle,
	ruleNorwayFirst,
	func(s Street) ([]Street, bool) { return neighborRule(s, MaskSmoke, Chesterfields, MaskPet, Fox) },
	func(s Street) ([]Street, bool) { return neighborRule(s, MaskSmoke, Kools, MaskPet, Horse) },
	func(s Street) ([]Street, bool) { return neighborRule(s, MaskNationality, Norwegian, MaskColor, Blue) },
}

func solver() Street {
	// Start with a random street.
	attribs := [][]uint16{
		{Englishman, Spaniard, Japanese, Ukrainian, Norwegian},
		{Dog, Fox, Snails, Horse, Zebra},
		{Red, Ivory, Green, Blue, Yellow},
		{Tea, Coffee, OrangeJuice, Milk, Water},
		{OldGold, Kools, LuckyStrike, Parliaments, Chesterfields},
	}
	for _, slice := range attribs {
		rand.Shuffle(5, func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	}
	street := Street{}
	for i := range 5 {
		for _, attrib := range attribs {
			street[i] += attrib[i]
		}
	}

	// Apply rules repeatedly and queue new mutations until we find a street that satisfies all the rules.
	todo := []Street{street}
	seen := make(map[Street]bool)
	for len(todo) != 0 {
		// Pop a street to use.
		street := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		found := true
		// Apply all rules. Either we satisfy them all or we craft better alternatives.
		for _, rule := range rules {
			if alternatives, ok := rule(street); !ok {
				for _, alternative := range alternatives {
					if seen[alternative] {
						continue
					}
					seen[alternative] = true
					todo = append(todo, alternative)
				}
				found = false
				break
			}
		}
		if found {
			return street
		}
	}
	return Street{}
}

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	street := solver()
	return Solution{
		DrinksWater: street.nationality(MaskDrink, Water),
		OwnsZebra:   street.nationality(MaskPet, Zebra),
	}
}
