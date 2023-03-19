package main

import "math/rand"

type (
	element struct {
		score int
		val int
		forward []*element
	}
	skiplist struct {
		header *element
		level int
		length int
	}
)

const (
	max_level = 16
	p float32 = 0.25
) 

func new_skiplist() *skiplist {
	return &skiplist{
		header: &element{forward:  make([]*element, 0)},
	}
}
func new_element(score, val int, level int) *element {
	return &element{
		score: score,
		val: val,
		forward: make([]*element, level),
	}
}

func (sk *skiplist) search(score int) *element {
	cur := sk.header
	for i := sk.level -1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score < score {
			cur = cur.forward[i]
		}
	}
	target := cur.forward[0]
	if target != nil && target.score != score {
		return nil
	}
	return target
}

func (sk *skiplist) insert(score, val int) (elm *element) {
	cur := sk.header
	update := make([]*element, max_level)
	for i := sk.level-1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score < score {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	level := random_level()
	if level > sk.level {
		level = sk.level + 1
		update[level] = sk.header
		sk.level = level
	}
	new_elem := new_element(score, val, level)
	for i := 0; i < sk.level; i++ {
		new_elem.forward[i] = update[i].forward[i]
		update[i].forward[i] = new_elem
	}
	sk.length++
	return new_elem
}

func (sk *skiplist) remove(score, val int)(elm *element, ok bool) {
	cur := sk.header
	update := make([]*element, max_level)
	for i := sk.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].score < score {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	target := cur.forward[0]
	if target == nil || target.score != score {
		return nil, false
	}
	for i := 0; i < sk.level; i++ {
		if update[i].forward[i] != target {
			// 理解上。。。
			// return nil, false 
			break;
		}
		update[i].forward[i] = target.forward[i]
	}
	sk.length--
	return target, true
}	

func random_level() int {
	level := 1
    for rand.Float32() < p && level < max_level {
        level++
    }
    return level
}