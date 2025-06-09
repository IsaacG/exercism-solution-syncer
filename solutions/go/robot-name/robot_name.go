// Package robotname names robots.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var used = make(map[string]bool)
var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getName() (string, error) {
	if len(used) == 26*26*10*10*10 {
		return "", errors.New("out of names")
	}
	var n string
	for {
		n = fmt.Sprintf(
			"%c%c%d%d%d",
			letters[rand.Intn(26)],
			letters[rand.Intn(26)],
			rand.Intn(10),
			rand.Intn(10),
			rand.Intn(10))
		if _, ok := used[n]; !ok {
			break
		}
	}
	used[n] = true
	return n, nil
}

// Robot is a named machine.
type Robot struct {
	name string
}

// Name returns the robot's name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		n, err := getName()
		if err != nil {
			return "", err
		}
		r.name = n
	}
	return r.name, nil
}

// Reset resets the robot's name.
func (r *Robot) Reset() {
	r.name = ""
}
