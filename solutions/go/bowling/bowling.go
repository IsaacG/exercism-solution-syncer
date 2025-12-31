package bowling

import (
	"errors"
	"slices"
)

const (
	maxPins = 10
	regularFrames = 9
	maxFrame = regularFrames+1
	maxRolls = 2
	maxRollsLastFrame = 3
)

// sum returns the sum of a slice of ints.
func sum(vals []int) int {
	var total int
	for _, v := range vals {
		total += v
	}
	return total
}

// frame stores one game frame.
type frame struct {
	num   int
	rolls []int
	next  *frame
}

// nextTwoRolls returns two more rolls for use in computing a strike or spare.
func (f *frame) nextTwoRolls() []int {
	rolls := slices.Clone(f.next.rolls)
	if len(rolls) > 2 {
		rolls = rolls[:2]
	} else if len(rolls) < 2 {
		rolls = append(rolls, f.next.next.rolls[0])
	}
	return rolls
}

// score returns the score of a frame, handling a strike or spare.
func (f *frame) score() int {
	if len(f.rolls) == 0 {
		return 0
	}
	total := sum(f.rolls)
	if f.num <= regularFrames {
		if f.rolls[0] == maxPins {
			total += sum(f.nextTwoRolls())
		} else if total == maxPins {
			total += f.nextTwoRolls()[0]
		}
	}
	return total
}

// full returns if a frame is full/complete.
func (f *frame) full() bool {
	if f.num <= regularFrames {
		return len(f.rolls) == maxRolls || (len(f.rolls) == maxRolls-1 && sum(f.rolls) == maxPins)
	}
	return len(f.rolls) == maxRollsLastFrame || (len(f.rolls) == maxRollsLastFrame-1 && sum(f.rolls) < maxPins)
}

// Game represents a bowling game.
type Game struct {
	firstFrame   *frame
	currentFrame *frame
	frameCount int
}

// NewGame returns a new Game with the first frame ready to use.
func NewGame() *Game {
	firstFrame := &frame{}
	return &Game{firstFrame: firstFrame, currentFrame: firstFrame, frameCount: 1}
}

// addFrameIfFull adds the next Frame to a game if the current frame is full.
func (g *Game) addFrameIfFull() {
	currentFrame := g.currentFrame
	if currentFrame.full() {
		g.frameCount++
		new := &frame{num: g.frameCount}
		currentFrame.next = new
		g.currentFrame = g.currentFrame.next
	}
}

// Roll adds a roll to the game.
func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > maxPins {
		return errors.New("invalid pins")
	}
	currentFrame := g.currentFrame
	var priorRoll int
	if len(currentFrame.rolls) > 0 {
		priorRoll = currentFrame.rolls[len(currentFrame.rolls)-1]
	}
	if g.frameCount <= regularFrames && len(currentFrame.rolls) > 0 && priorRoll != maxPins && priorRoll+pins > maxPins {
		return errors.New("invalid frame")
	}
	if g.frameCount == maxFrame && sum(currentFrame.rolls) != maxPins && priorRoll != maxPins && priorRoll+pins > maxPins {
		return errors.New("invalid frame")
	}
	if g.frameCount > maxFrame {
		return errors.New("invalid game over")
	}
	currentFrame.rolls = append(currentFrame.rolls, pins)
	g.addFrameIfFull()

	return nil
}

// Score returns the score of a game.
func (g *Game) Score() (int, error) {
	if g.frameCount != maxFrame+1 {
		return 0, errors.New("game not over")
	}
	var total int
	for f := g.firstFrame; f.next != nil; f = f.next {
		total += f.score()
	}
	return total, nil
}
