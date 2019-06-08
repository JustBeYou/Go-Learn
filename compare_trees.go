package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, level int) {
	if t.Left != nil {
		Walk(t.Left, ch, level + 1)	
	} 
	if t.Right != nil {
		Walk(t.Right, ch, level + 1)
	}
	
	ch <- t.Value
	
	if level == 0 {
		close(ch)	
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	m := make(map [int] bool)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1, 0)
	go Walk(t2, ch2, 0)
	
	for {
		select {
			case val, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else if _, exists := m[val]; exists {
					m[val] = false;
				} else if ok {
					m[val] = true;	
				}
			
			case val, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else if _, exists := m[val]; exists {
					m[val] = false;
				} else if ok {
					m[val] = true;	
				}
		}
		
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	
	for _, v := range m {
		if v == true {
			return false	
		}
	}
	
	return true	
}
func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

