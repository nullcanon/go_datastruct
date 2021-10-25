package binary_tree

import (
	"data_struct/utils"
)

type Node struct {
	left  *Node
	right *Node
	Value interface{}
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

type Tree struct {
	root *Node
	size uint
}

func New() *Tree {
	Tree := &Tree{}
	return Tree
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Size() uint {
	return t.size
}

func (t *Tree) Insert(value interface{}, compare utils.Compare) {
	node := &Node{Value: value}
	if t.root == nil {
		t.root = node
	} else {
		pre := t.root
		for cur := t.root; cur != nil; {
			pre = cur
			if compare(value, cur.Value) <= 0 {
				cur = cur.left
			} else {
				cur = cur.right
			}
		}
		if compare(value, pre.Value) <= 0 {
			pre.left = node
		} else {
			pre.right = node
		}
	}
	t.size++
}

func (t *Tree) InsertR(value interface{}, compare utils.Compare) {
	if t.root == nil {
		t.root = &Node{Value: value}
		t.size++
		return
	}

	t.insertR(t.root, value, compare)
}

func (t *Tree) insertR(root *Node, value interface{}, compare utils.Compare) *Node {
	if root == nil {
		node := &Node{Value: value}
		t.size++
		return node
	}

	if compare(value, root.Value) <= 0 {
		root.left = t.insertR(root.left, value, compare)
	} else {
		root.right = t.insertR(root.right, value, compare)
	}
	return root
}

func (t *Tree) find(root *Node, value interface{}, compare utils.Compare) *Node {

	if root == nil {
		return nil
	}

	if compare(value, root.Value) < 0 {
		return t.find(root.left, value, compare)
	} else if compare(value, root.Value) > 0 {
		return t.find(root.right, value, compare)
	} else {
		return root
	}
}

func (t *Tree) Find(value interface{}, compare utils.Compare) *Node {
	return t.find(t.root, value, compare)
}

func (t *Tree) findParent(root *Node, parent *Node, value interface{}, compare utils.Compare) (*Node, *Node) {

	if root == nil {
		return nil, nil
	}

	if compare(value, root.Value) < 0 {
		return t.findParent(root.left, root, value, compare)
	} else if compare(value, root.Value) > 0 {
		return t.findParent(root.right, root, value, compare)
	} else {
		return root, parent
	}
}

func (t *Tree) FindParent(value interface{}, compare utils.Compare) (*Node, *Node) {
	return t.findParent(t.root, nil, value, compare)
}

func (t *Tree) Remove(value interface{}, compare utils.Compare) {

	if t.root == nil {
		return
	}

	del, parent := t.FindParent(value, compare)

	if del == nil {
		return
	}

	var root *Node
	var root_pre *Node

	// 要删除节点的左子树的最右节点
	for cur := del.left; cur != nil; cur = cur.right {

		root = cur
		if cur.right != nil && cur.right.right == nil {
			root_pre = cur
		}
	}
	// 左子树不为空
	if root != nil {
		if parent != nil {
			if parent.left == del {
				parent.left = root
			} else {
				parent.right = root
			}
		} else {
			t.root = root
		}
		if del.left != root && del.right != root{
			root.left = del.left
			root.right = del.right
		}
		if root_pre != nil {
			root_pre.right = nil
		}
	} else {
		// 左子树为空,在右子树找
		for cur := del.right; cur != nil; cur = cur.left {
			root = cur
			if cur.left != nil && cur.left.left == nil {
				root_pre = cur
			}
		}
		if root != nil {
			if parent != nil {
				if parent.left == del {
					parent.left = root
				} else {
					parent.right = root
				}
			} else {
				t.root = root
			}
			if del.right != root {
				root.left = del.left
				root.right = del.right
			}
			if root_pre != nil {
				root_pre.left = nil
			}
		} else {
			// 右子树也为空, 说明删除的是叶子节点
			if parent != nil {
				if parent.left == del {
					parent.left = nil
				} else {
					parent.right = nil
				}
			} else {
				t.root = nil
			}
		}
	}
	t.size--
}

func (t *Tree) preTraverseR(root *Node, visitor utils.Visitor) {
	if root != nil {
		visitor(root.Value)
		t.preTraverseR(root.left, visitor)
		t.preTraverseR(root.right, visitor)
	}
}

func (t *Tree) PreTraverseR(visitor utils.Visitor) {
	t.preTraverseR(t.root, visitor)
}

func (t *Tree) inTraverseR(root *Node, visitor utils.Visitor) {
	if root != nil {
		t.inTraverseR(root.left, visitor)
		visitor(root.Value)
		t.inTraverseR(root.right, visitor)
	}
}

func (t *Tree) InTraverseR(visitor utils.Visitor) {
	t.inTraverseR(t.root, visitor)
}

func (t *Tree) postTraverseR(root *Node, visitor utils.Visitor) {
	if root != nil {
		t.postTraverseR(root.left, visitor)
		t.postTraverseR(root.right, visitor)
		visitor(root.Value)
	}
}

func (t *Tree) PostTraverseR(visitor utils.Visitor) {
	t.postTraverseR(t.root, visitor)
}
