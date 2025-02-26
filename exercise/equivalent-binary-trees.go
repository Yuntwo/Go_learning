package exercise

import "golang.org/x/tour/tree"

// 时间复杂度遍历一次n个节点，O(n)
// 空间复杂度递归压栈n个节点，O(n)

// 这样是两个channel并发，其实直接同步的话时间和空间复杂度只会是2倍而已

// TODO 不太清楚有没有更优的channel并发方式

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkReg(t, ch)
	close(ch)
}

func Walk2(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk2(t.Left, ch)
	ch <- t.Value
	Walk2(t.Right, ch)
}

func WalkReg(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// 注意这里都要调用WalkReg，而不是Walk
	WalkReg(t.Left, ch)
	ch <- t.Value
	WalkReg(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	// 也可以直接这样用匿名函数写
	//go func() {
	//    Walk2(t1, ch1)
	//    close(ch1)
	//} ()
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		if !ok1 && !ok2 {
			return true
		}
		if (!ok1 && ok2) || (ok1 && !ok2) || (x != y) {
			return false
		}
	}
	return true
}

func main() {
	Same(tree.New(2), tree.New(2))
}
