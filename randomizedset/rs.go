package randomizedset

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	d         map[int]int
	positions []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		d:         map[int]int{},
		positions: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.d[val]
	if ok {
		return false
	}

	this.positions = append(this.positions, val)

	this.d[val] = len(this.positions) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.d[val]
	if !ok {
		return ok
	}

	this.positions = append(this.positions[:v], this.positions[v+1:]...)

	for i := v; i < len(this.positions); i++ {
		this.d[this.positions[i]] = i
	}

	delete(this.d, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.d) == 0 {
		return 0
	}
	if len(this.d) == 1 {
		return this.positions[0]
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(this.positions))
	if n == len(this.positions) {
		n--
	}
	return this.positions[n]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
