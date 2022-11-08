package searching

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	newBst := func() *bst[intWrapper]{
		return &bst[intWrapper]{}
	}

	t.Run("empty", func(t *testing.T) {
		b := newBst()
		assert.Equal(t, []intWrapper{}, b.collect())
		
		_, ok := b.max()
		assert.False(t, ok)
		
		_, ok = b.min()
		assert.False(t, ok)
	})

	t.Run("multiple", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		assert.Equal(t, []intWrapper{1,2,3,5,7}, b.collect())
		
		max, ok := b.max()
		assert.True(t, ok)
		assert.Equal(t, 7, int(max))

		min, ok := b.min()
		assert.True(t, ok)
		assert.Equal(t, 1, int(min))
	})
		
	t.Run("min max interleaved", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		max, ok := b.max()
		assert.True(t, ok)
		assert.Equal(t, 7, int(max))

		min, ok := b.min()
		assert.True(t, ok)
		assert.Equal(t, 1, int(min))

		for _, v := range []intWrapper{88,-123,32,12,-5} {
			b.add(v)
		}

		max, ok = b.max()
		assert.True(t, ok)
		assert.Equal(t, 88, int(max))

		min, ok = b.min()
		assert.True(t, ok)
		assert.Equal(t, -123, int(min))
	})

	t.Run("del min", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		b.delMin()
		assert.Equal(t, []intWrapper{2,3,5,7}, b.collect())
	})

	t.Run("del min 2", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{88,-123,32,12,-5} {
			b.add(v)
		}
		b.delMin()
		assert.Equal(t, []intWrapper{-5,12,32,88}, b.collect())
	})

	t.Run("get", func(t *testing.T) {
		b := newBst()
		for _, v := range []intWrapper{7,5,3,2,5,7,1,2} {
			b.add(v)
		}
		got, ok := b.get(5)
		assert.True(t, ok)
		assert.Equal(t, got, intWrapper(5))

		_, ok = b.get(125)
		assert.False(t, ok)

		got, ok = b.getRec(5)
		assert.True(t, ok)
		assert.Equal(t, got, intWrapper(5))

		_, ok = b.getRec(125)
		assert.False(t, ok)
	})
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		a int
		b int
		exp int		
	}{
		{2,5,-1},
		{5,2,1},
		{5,5,0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v %v", tC.a, tC.b), func(t *testing.T) {
			assert.Equal(t, tC.exp, intWrapper(tC.a).cmp(intWrapper(tC.b)))
		})
	}
}

func TestBstDelete(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.delete(123)
		b.delete(1)
		b.delete(3)

		assert.Equal(t, []intWrapper{}, b.collect())
	})

	t.Run("single el", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(123)
		b.delete(123)

		assert.Equal(t, []intWrapper{}, b.collect())
	})

	t.Run("root when 1 left child", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(1)
		b.delete(5)

		assert.Equal(t, []intWrapper{1}, b.collect())
	})

	t.Run("root when 1 right child", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(10)
		b.delete(5)

		assert.Equal(t, []intWrapper{10}, b.collect())
	})

	t.Run("root when 2 children", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(5)
		b.add(10)
		b.add(1)
		b.delete(5)

		assert.Equal(t, []intWrapper{1,10}, b.collect())
	})

	t.Run("not empty", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}
		b.delete(123)
		b.delete(1)
		b.delete(7)
		b.delete(5)

		assert.Equal(t, []intWrapper{3,4,8}, b.collect())
	})

	t.Run("root when many", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}
		b.delete(5)

		assert.Equal(t, []intWrapper{1,3,4,7,8}, b.collect())
	})

	t.Run("tall tree", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{1,2,3,10,11,12, 5,6,7} {
			b.add(intWrapper(v))
		}
		b.delete(10)
		b.delete(11)

		assert.Equal(t, []intWrapper{1,2,3,5,6,7,12}, b.collect())
	})

	t.Run("interleaved", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []int{5,3,7,4,1,8} {
			b.add(intWrapper(v))
		}

		b.delete(1)
		b.delete(7)
		b.delete(5)

		for _, v := range []int{2,10,-14,1} {
			b.add(intWrapper(v))
		}

		assert.Equal(t, []intWrapper{-14,1,2,3,4,8,10}, b.collect())
	})
}

func TestBstRange(t *testing.T) {
	testCases := []struct {
		min int
		max int
		exp []intWrapper
	}{
		{ 1,10,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ 1,5,[]intWrapper{1,2,3,4,5} },
		{ 1,100,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ -100,100,[]intWrapper{1,2,3,4,5,6,7,8,9,10} },
		{ -100,3,[]intWrapper{1,2,3} },
		{ 5,10,[]intWrapper{5,6,7,8,9,10} },
		{ 10,8, []intWrapper{}},
		{ 1,1,[]intWrapper{1}},
		{ 1,2, []intWrapper{1,2}},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tC.min, tC.max), func(t *testing.T) {
			b := &bst[intWrapper]{}
			// 1,2,3,4,5,6,7,8,9,10
			for _, v := range []intWrapper{4,10,9,2,5,6,3,8,1,7} {
				b.add(v)
			}
			assert.Equal(t, tC.exp, b.inRange(intWrapper(tC.min),intWrapper(tC.max)))
		})
	}
}

func TestHeight(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := &bst[intWrapper]{}
		assert.Equal(t, 0, b.height())
	})

	t.Run("single", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(3)
		assert.Equal(t, 1, b.height())
	})

	t.Run("2 elements", func(t *testing.T) {
		b := &bst[intWrapper]{}
		b.add(1)
		b.add(3)
		assert.Equal(t, 2, b.height())
	})

	t.Run("right tall tree", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []intWrapper{1,2,3,4,5,6,7,8} {
			b.add(v)
		}
		assert.Equal(t, 8, b.height())
	})

	t.Run("left tall tree", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []intWrapper{8,7,6,5,4,3,2,1} {
			b.add(v)
		}
		assert.Equal(t, 8, b.height())
	})

	t.Run("tall tree but with leaf", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []intWrapper{2,3,4,5,6,7,8,1} {
			b.add(v)
		}
		assert.Equal(t, 7, b.height())
	})

	t.Run("mixed", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []intWrapper{4,10,9,2,5,6,3,8,1,7} {
			b.add(v)
		}
		assert.Equal(t, 7, b.height())
	})
	
	t.Run("mixed2", func(t *testing.T) {
		b := &bst[intWrapper]{}
		for _, v := range []intWrapper{8,3,10,1,6,14,4,7,13} {
			b.add(v)
		}
		assert.Equal(t, 4, b.height())
	})
}

func TestBfs(t *testing.T) {
	t.Fatal("todo")
}