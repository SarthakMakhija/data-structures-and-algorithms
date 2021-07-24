package binarytree_test

import (
	binarytree "github.com/SarthakMakhija/data-structures-and-algorithms/tree/binary"
	"reflect"
	"testing"
)

func TestExistsRootToLeafPathWithSum1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 314,
		Left: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 271,
				Left: &binarytree.IntNode{
					Value: 28,
				},
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
					Left: &binarytree.IntNode{
						Value: 17,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 401,
						Right: &binarytree.IntNode{
							Value: 641,
						},
					},
					Right: &binarytree.IntNode{
						Value: 257,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 271,
				Right: &binarytree.IntNode{
					Value: 28,
				},
			},
		},
	}

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	exists := tree.ExistsRootToLeafPathWith(591)
	if exists != true {
		t.Fatalf("Expected true, received %v", exists)
	}
}

func TestExistsRootToLeafPathWithSum2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 314,
		Left: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 271,
				Left: &binarytree.IntNode{
					Value: 28,
				},
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
					Left: &binarytree.IntNode{
						Value: 17,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 401,
						Right: &binarytree.IntNode{
							Value: 641,
						},
					},
					Right: &binarytree.IntNode{
						Value: 257,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 271,
				Right: &binarytree.IntNode{
					Value: 28,
				},
			},
		},
	}

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	exists := tree.ExistsRootToLeafPathWith(590)
	if exists != false {
		t.Fatalf("Expected false, received %v", exists)
	}
}

func TestAllPathsFromRootToLeafWithSum1(t *testing.T) {
	root := binarytree.IntNode{
		Value: 314,
		Left: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 271,
				Left: &binarytree.IntNode{
					Value: 28,
				},
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
					Left: &binarytree.IntNode{
						Value: 17,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 401,
						Right: &binarytree.IntNode{
							Value: 641,
						},
					},
					Right: &binarytree.IntNode{
						Value: 257,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 271,
				Right: &binarytree.IntNode{
					Value: 28,
				},
			},
		},
	}

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}

	paths := tree.AllPathsFromRootToLeafWith(591)

	values := paths[0].NodeValues
	expected := []int{314, 6, 271, 0}

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("Expected paths %v, received %v", expected, paths)
	}
}

func TestAllPathsFromRootToLeafWithSum2(t *testing.T) {
	root := binarytree.IntNode{
		Value: 314,
		Left: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 271,
				Left: &binarytree.IntNode{
					Value: 28,
				},
				Right: &binarytree.IntNode{
					Value: 0,
				},
			},
			Right: &binarytree.IntNode{
				Value: 561,
				Right: &binarytree.IntNode{
					Value: 3,
					Left: &binarytree.IntNode{
						Value: 17,
					},
				},
			},
		},
		Right: &binarytree.IntNode{
			Value: 6,
			Left: &binarytree.IntNode{
				Value: 2,
				Right: &binarytree.IntNode{
					Value: 1,
					Left: &binarytree.IntNode{
						Value: 401,
						Right: &binarytree.IntNode{
							Value: 641,
						},
					},
					Right: &binarytree.IntNode{
						Value: 257,
					},
				},
			},
			Right: &binarytree.IntNode{
				Value: 271,
				Right: &binarytree.IntNode{
					Value: 28,
				},
			},
		},
	}

	tree := binarytree.IntBinaryTree{
		Root: &root,
	}
	paths := tree.AllPathsFromRootToLeafWith(619)
	expected := []binarytree.Path{
		{
			NodeValues: []int{314, 6, 271, 28},
		},
		{
			NodeValues: []int{314, 6, 271, 28},
		},
	}

	if !reflect.DeepEqual(paths, expected) {
		t.Fatalf("Expected paths %v, received %v", expected, paths)
	}
}
