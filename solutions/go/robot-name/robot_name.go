// Package robotname names robots.
package robotname

import (
	"errors"
	"math/rand"
	"strconv"
)

var used = make(map[string]bool)
var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getName() (string, error) {
	if len(used) == 26*26*10*10*10 {
		return "", errors.New("out of names")
	}
	name := make([]byte, 0, 5)
	var n string
	for {
		for i := 0; i < 2; i++ {
			name = append(name, letters[rand.Intn(26)])
		}
		for i := 0; i < 3; i++ {
			name = append(name, strconv.Itoa(rand.Intn(10))[0])
		}
		n = string(name)
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
