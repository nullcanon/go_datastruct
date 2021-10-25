package binary_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	tree := New()
	assert.Equal(t, uint(0), tree.Size())
}

func TestInsert(t *testing.T) {
	tree := New()
	assert.Equal(t, uint(0), tree.Size())

	cmp := func(left interface{}, right interface{}) int {
		return left.(int) - right.(int)
	}

	tree.Insert(10, cmp)
	assert.Equal(t, uint(1), tree.Size())
	assert.Equal(t, tree.Root().Value, 10)

	tree.Insert(20, cmp)
	assert.Equal(t, uint(2), tree.Size())
	assert.Equal(t, tree.Root().Right().Value, 20)

	tree.Insert(5, cmp)
	assert.Equal(t, uint(3), tree.Size())
	assert.Equal(t, tree.Root().Left().Value, 5)

	tree.Insert(6, cmp)
	assert.Equal(t, uint(4), tree.Size())
	assert.Equal(t, tree.Root().Left().Right().Value, 6)

}

func TestInsertR(t *testing.T) {
	tree := New()
	assert.Equal(t, uint(0), tree.Size())

	cmp := func(left interface{}, right interface{}) int {
		return left.(int) - right.(int)
	}

	tree.InsertR(10, cmp)
	assert.Equal(t, uint(1), tree.Size())
	assert.Equal(t, tree.Root().Value, 10)

	tree.InsertR(20, cmp)
	assert.Equal(t, uint(2), tree.Size())
	assert.Equal(t, tree.Root().Right().Value, 20)

	tree.InsertR(5, cmp)
	assert.Equal(t, uint(3), tree.Size())
	assert.Equal(t, tree.Root().Left().Value, 5)

	tree.InsertR(6, cmp)
	assert.Equal(t, uint(4), tree.Size())
	assert.Equal(t, tree.Root().Left().Right().Value, 6)
}

func TestTraverseRAndFind(t *testing.T) {
	tree := New()
	assert.Equal(t, uint(0), tree.Size())

	cmp := func(left interface{}, right interface{}) int {
		return left.(int) - right.(int)
	}

	ret := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret = append(ret, value.(int))
		return true
	})
	assert.Equal(t, "[]", fmt.Sprintf("%v", ret))

	ret5 := make([]int, 0)
	tree.InTraverseR(func(value interface{}) bool {
		ret5 = append(ret5, value.(int))
		return true
	})
	assert.Equal(t, "[]", fmt.Sprintf("%v", ret5))

	ret6 := make([]int, 0)
	tree.PostTraverseR(func(value interface{}) bool {
		ret6 = append(ret6, value.(int))
		return true
	})

	assert.Equal(t, "[]", fmt.Sprintf("%v", ret6))

	tree.InsertR(6, cmp)
	tree.InsertR(2, cmp)
	tree.InsertR(4, cmp)
	tree.InsertR(8, cmp)
	tree.InsertR(10, cmp)
	tree.InsertR(7, cmp)
	tree.InsertR(1, cmp)
	assert.Equal(t, uint(7), tree.Size())
	/*
			    6
			2       8
		  1   4   7   10
	*/

	ret2 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret2 = append(ret2, value.(int))
		return true
	})
	assert.Equal(t, "[6 2 1 4 8 7 10]", fmt.Sprintf("%v", ret2))

	ret3 := make([]int, 0)
	tree.InTraverseR(func(value interface{}) bool {
		ret3 = append(ret3, value.(int))
		return true
	})
	assert.Equal(t, "[1 2 4 6 7 8 10]", fmt.Sprintf("%v", ret3))

	ret4 := make([]int, 0)
	tree.PostTraverseR(func(value interface{}) bool {
		ret4 = append(ret4, value.(int))
		return true
	})
	assert.Equal(t, "[1 4 2 7 10 8 6]", fmt.Sprintf("%v", ret4))

	assert.Equal(t, 6, tree.Find(6, cmp).Value)
	assert.Equal(t, 1, tree.Find(1, cmp).Value)
	assert.Equal(t, (*Node)(nil), tree.Find(100, cmp))

	cur, par := tree.FindParent(6, cmp)
	assert.Equal(t, cur.Value, 6)
	assert.Equal(t, par, (*Node)(nil))

	cur, par = tree.FindParent(1, cmp)
	assert.Equal(t, cur.Value, 1)
	assert.Equal(t, par.Value, 2)
}

func TestRemove(t *testing.T) {
	tree := New()
	assert.Equal(t, uint(0), tree.Size())

	cmp := func(left interface{}, right interface{}) int {
		return left.(int) - right.(int)
	}

	tree.InsertR(6, cmp)
	tree.InsertR(2, cmp)
	tree.InsertR(4, cmp)
	tree.InsertR(8, cmp)
	tree.InsertR(10, cmp)
	tree.InsertR(7, cmp)
	tree.InsertR(1, cmp)

	/*
			    6
			2       8
		  1   4   7   10
	*/

	tree.Remove(6, cmp)
	assert.Equal(t, tree.Root().Value, 4)
	ret2 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret2 = append(ret2, value.(int))
		return true
	})
	assert.Equal(t, "[4 2 1 8 7 10]", fmt.Sprintf("%v", ret2))

	/*
			    4
			2       8
		  1       7   10
	*/
	tree.Remove(2, cmp)
	assert.Equal(t, tree.Root().Left().Value, 1)
	ret3 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret3 = append(ret3, value.(int))
		return true
	})
	assert.Equal(t, "[4 1 8 7 10]", fmt.Sprintf("%v", ret3))

	/*
			    4
			1       8
		          7   10
	*/
	tree.Remove(1, cmp)
	assert.Equal(t, tree.Root().Left(), (*Node)(nil))
	ret4 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret4 = append(ret4, value.(int))
		return true
	})
	assert.Equal(t, "[4 8 7 10]", fmt.Sprintf("%v", ret4))


	
	/*
			    4
			        8
		          7   10
	*/
	fmt.Print("--------------------------------")
	tree.Remove(7, cmp)
	ret5 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret5 = append(ret5, value.(int))
		return true
	})
	assert.Equal(t, "[4 8 10]", fmt.Sprintf("%v", ret5))


		/*
			    4
			        8
		              10
	*/
	tree.Remove(4, cmp)
	ret6 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret6 = append(ret6, value.(int))
		return true
	})
	assert.Equal(t, "[8 10]", fmt.Sprintf("%v", ret6))


	
		/*
			    
			        8
		              10
	*/
	tree.Remove(10, cmp)
	ret7 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret7 = append(ret7, value.(int))
		return true
	})
	assert.Equal(t, "[8]", fmt.Sprintf("%v", ret7))

			/*
			    
			        
		              8
	*/
	tree.Remove(8, cmp)
	ret8 := make([]int, 0)
	tree.PreTraverseR(func(value interface{}) bool {
		ret8 = append(ret8, value.(int))
		return true
	})
	assert.Equal(t, "[]", fmt.Sprintf("%v", ret8))
}
