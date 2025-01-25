package leetcode636

import (
	"fmt"
	"strconv"
	"strings"
)

//#array #stack
// https://leetcode.com/problems/exclusive-time-of-functions/

func exclusiveTime(n int, logs []string) []int {
	stack := NewBackwardStack()
	res := make([]int, n)
	prevTime := 0

	for _, v := range logs {

		var id, time int
		var status string

		_s := strings.Split(v, ":")
		status = _s[1]
		id, _ = strconv.Atoi(_s[0])
		time, _ = strconv.Atoi(_s[2])

		// fmt.Println(id, status, time, i)

		if status == "start" {
			if stack.Len() > 0 {
				res[stack.Current()] += time - prevTime
			}
			stack.Push(id)
			prevTime = time
		} else {
			fid, _ := stack.Pop()
			res[fid] += time - prevTime + 1
			prevTime = time + 1
		}

	}

	return res
}

func NewBackwardStack() *BackwardStack {
	return &BackwardStack{
		stack:     []int{},
		currentid: 0,
	}
}

type BackwardStack struct {
	stack     []int
	currentid int
}

func (b *BackwardStack) Current() int {
	if b.currentid < 0 {
		return -1
	}
	return b.stack[b.currentid]
}

func (b *BackwardStack) ByID(id int) int {
	if id < 0 || id > b.currentid {
		return -1
	}
	return b.stack[id]
}

func (b *BackwardStack) Len() int {
	if b.currentid < 0 {
		return 0
	}
	return len(b.stack) - b.currentid
}

func (b *BackwardStack) Push(value int) {
	if len(b.stack) == 0 {
		b.stack = append(b.stack, value)
		b.currentid = 0
		return
	}

	if b.currentid+1 < cap(b.stack) {

		if b.currentid+1 < len(b.stack) {
			b.stack[b.currentid+1] = value
			b.currentid = b.currentid + 1
			return
		}

		if b.currentid+1 >= len(b.stack) {
			b.stack = append(b.stack, value)
			b.currentid = len(b.stack) - 1
			return
		}

		return
	}

	if b.currentid+1 >= cap(b.stack) {

		b.stack = append(b.stack, value)
		b.currentid = len(b.stack) - 1

		return
	}
}

func (b *BackwardStack) Pop() (int, bool) {
	if len(b.stack) == 0 {
		return 0, false
	}
	if b.currentid < 0 {
		return 0, false
	}

	retval := b.stack[b.currentid]
	b.currentid = b.currentid - 1

	return retval, true
}

func RunExclusiveTime() {
	n := 2
	logs := []string{"0:start:0", "0:end:0", "1:start:1", "1:end:1", "2:start:2", "2:end:2", "2:start:3", "2:end:3"}
	fmt.Println(exclusiveTime(n, logs))

	// stack := NewBackwardStack()
	// stack.Push(1)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Len())
	// fmt.Println(stack.Current())
	// stack.Push(2)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Len())
	// fmt.Println(stack.Current())
	// stack.Push(3)
	// fmt.Println(stack.Pop())
	// fmt.Println(stack.Len())
	// fmt.Println(stack.Current())

}

func exclusiveTime2(n int, logs []string) []int {
	stack := []int{}
	times := make([]int, n)
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		fid, _ := strconv.Atoi(parts[0])
		typ := parts[1]
		time, _ := strconv.Atoi(parts[2])

		if typ == "start" {
			if len(stack) > 0 {
				times[stack[len(stack)-1]] += time - prevTime
			}
			stack = append(stack, fid)
			prevTime = time
		} else {
			times[stack[len(stack)-1]] += time - prevTime + 1
			stack = stack[:len(stack)-1]
			prevTime = time + 1
		}
	}

	return times
}
