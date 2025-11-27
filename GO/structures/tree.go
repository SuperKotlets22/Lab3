package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CompleteTree struct {
	Root *TreeNode
}

func NewTree() *CompleteTree {
	return &CompleteTree{}
}

func (t *CompleteTree) Insert(val int) {
	newNode := &TreeNode{Val: val}
	if t.Root == nil {
		t.Root = newNode
		return
	}
	// bfs вставка
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Left == nil {
			curr.Left = newNode
			return
		} else {
			queue = append(queue, curr.Left)
		}
		if curr.Right == nil {
			curr.Right = newNode
			return
		} else {
			queue = append(queue, curr.Right)
		}
	}
}

func (t *CompleteTree) Find(val int) bool {
	if t.Root == nil { return false }
	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Val == val { return true }
		if curr.Left != nil { queue = append(queue, curr.Left) }
		if curr.Right != nil { queue = append(queue, curr.Right) }
	}
	return false
}
