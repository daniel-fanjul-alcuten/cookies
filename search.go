package main

import (
	"container/list"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Stock float64
type Production float64
type Time float64
type Acceleration float64

func (s Stock) Divide(p Production) Time {
	return Time(float64(s) / float64(p))
}

func (p Production) Divide(t Time) Acceleration {
	return Acceleration(float64(p) / float64(t))
}

func (t Time) String() string {
	const MaxInt64 = 9223372036854775807
	s := t * Time(time.Second)
	if s > MaxInt64 {
		return "?"
	}
	return time.Duration(s).String()
}

type UpdateName string
type UpdateDescription string

type State interface {
	Updates() []Update
	Production() Production
}

type Update struct {
	State
	Name        UpdateName
	Description UpdateDescription
	Cost        Stock
}

type Node struct {
	State
	Production  Production
	Transitions TransitionMap
	Best        TransitionPath
}

type Transition struct {
	*Node
	Cost        Stock
	Time        Time
	Description UpdateDescription
}
type TransitionMap map[UpdateName]*Transition
type TransitionPath []*Transition

func (n *Node) ExpandTransitions() TransitionMap {
	if n.Transitions == nil {
		updates := n.Updates()
		n.Transitions = make(TransitionMap, len(updates))
		for _, update := range updates {
			node := &Node{update.State, update.Production(), nil, nil}
			time := update.Cost.Divide(n.Production)
			transition := &Transition{node, update.Cost, time, update.Description}
			n.Transitions[update.Name] = transition
		}
	}
	return n.Transitions
}

func (n *Node) ExpandBest(depth int) TransitionPath {
	if depth > len(n.Best) {
		type b struct {
			transition   *Transition
			acceleration Acceleration
		}
		var best b
		for _, transition := range n.ExpandTransitions() {
			time, production := transition.Time, transition.Production
			acceleration := (production - n.Production).Divide(time)
			if best.transition == nil || acceleration > best.acceleration {
				best = b{transition, acceleration}
			}
			if depth > 1 {
				tail := transition.ExpandBest(depth - 1)
				for _, transition := range tail {
					time += transition.Time
					production = transition.Production
					acceleration := (production - n.Production).Divide(time)
					if best.transition == nil || acceleration > best.acceleration {
						best = b{transition, acceleration}
					}
				}
			}
		}
		n.Best = append([]*Transition{best.transition},
			best.transition.ExpandBest(depth-1)...)
	}
	if depth > 0 {
		return n.Best[:depth]
	}
	return nil
}

type GoalStack struct {
	list.List
}

type ExpandBestGoal struct {
	*Node
	depth int
}

func (s *GoalStack) Tick(limit int) {
	ticked := 0
	for ticked < limit && s.Len() > 0 {
		value := s.Remove(s.Back())
		switch g := value.(type) {
		case ExpandBestGoal:
			if g.depth > len(g.Best) {
				ok := true
				if g.depth > 1 {
					for _, t := range g.ExpandTransitions() {
						if len(g.Best) > len(t.Best) {
							s.PushBack(ExpandBestGoal{t.Node, len(t.Best) + 1})
							ok = false
						}
					}
					if len(g.Best) > 0 {
						t := g.Best[0]
						if len(g.Best) > len(t.Best) {
							s.PushBack(ExpandBestGoal{t.Node, len(t.Best) + 1})
						}
					}
				}
				if ok {
					g.ExpandBest(len(g.Best) + 1)
				}
			}
		}
		ticked++
	}
}

func Search(state State, input <-chan string, output chan<- *Node, depth, number, limit int) {
	node, stack := &Node{state, state.Production(), nil, nil}, GoalStack{}
	parseInput := func(str string) {
		if len(str) == 0 {
			if len(node.Best) > 0 {
				transition := node.Best[0]
				node, stack = transition.Node, GoalStack{}
			}
		} else if str == "0" {
			node.Transitions = nil
			node.Best = nil
			stack = GoalStack{}
		} else if str == "gc" {
			runtime.GC()
		} else if str == "d" {
			fmt.Printf("d=%v\n", depth)
		} else if str == "n" {
			fmt.Printf("n=%v\n", number)
		} else if str == "l" {
			fmt.Printf("l=%v\n", limit)
		} else if strings.HasPrefix(str, "d=") {
			if i, err := strconv.Atoi(str[2:]); err == nil {
				depth = i
			}
		} else if strings.HasPrefix(str, "n=") {
			if i, err := strconv.Atoi(str[2:]); err == nil {
				number = i
			}
		} else if strings.HasPrefix(str, "l=") {
			if i, err := strconv.Atoi(str[2:]); err == nil {
				limit = i
			}
		} else {
			var count int
			var str2 string
			n, err := fmt.Sscanf(str, "%d%s", &count, &str2)
			if n == 2 && err == nil {
				str = str2
			} else {
				count = 1
			}
			name := UpdateName(str)
			for i := 0; i < count; i++ {
				transition := node.ExpandTransitions()[name]
				if transition == nil {
					break
				}
				node, stack = transition.Node, GoalStack{}
			}
		}
	}
	output <- node
	for {
		if stack.Len() == 0 {
			i, next := 0, node
			for i < number && depth <= len(next.Best) {
				i, next = i+1, next.Best[0].Node
			}
			if depth > len(next.Best) {
				stack.PushBack(ExpandBestGoal{next, depth})
			}
		}
		if stack.Len() == 0 {
			str, ok := <-input
			if !ok {
				close(output)
				return
			}
			parseInput(str)
		} else {
			select {
			case str, ok := <-input:
				if !ok {
					close(output)
					return
				}
				parseInput(str)
			default:
				stack.Tick(limit)
			}
		}
		select {
		case output <- node:
		default:
		}
	}
}
