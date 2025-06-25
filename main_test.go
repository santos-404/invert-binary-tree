package main

import (
	"reflect"
	"testing"
)

func NewNode(val int, left, right *BinaryTreeNode) *BinaryTreeNode {
	return &BinaryTreeNode{
		Value:      val,
		LeftChild:  left,
		RightChild: right,
	}
}

func TestInvertBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		input    *BinaryTreeNode
		expected *BinaryTreeNode
	}{
		{
			name:     "nil tree",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single node",
			input:    NewNode(1, nil, nil),
			expected: NewNode(1, nil, nil),
		},
		{
			name: "left-heavy tree",
			input: NewNode(1,
				NewNode(2,
					NewNode(3, nil, nil),
					nil,
				),
				nil,
			),
			expected: NewNode(1,
				nil,
				NewNode(2,
					nil,
					NewNode(3, nil, nil),
				),
			),
		},
		{
			name: "right-heavy tree",
			input: NewNode(1,
				nil,
				NewNode(2,
					nil,
					NewNode(3, nil, nil),
				),
			),
			expected: NewNode(1,
				NewNode(2,
					NewNode(3, nil, nil),
					nil,
				),
				nil,
			),
		},
		{
			name: "full binary tree",
			input: NewNode(1,
				NewNode(2,
					NewNode(4, nil, nil),
					NewNode(5, nil, nil),
				),
				NewNode(3,
					NewNode(6, nil, nil),
					NewNode(7, nil, nil),
				),
			),
			expected: NewNode(1,
				NewNode(3,
					NewNode(7, nil, nil),
					NewNode(6, nil, nil),
				),
				NewNode(2,
					NewNode(5, nil, nil),
					NewNode(4, nil, nil),
				),
			),
		},
		{
			name: "tree with some nil children",
			input: NewNode(1,
				NewNode(2,
					nil,
					NewNode(4, nil, nil),
				),
				NewNode(3,
					nil,
					nil,
				),
			),
			expected: NewNode(1,
				NewNode(3,
					nil,
					nil,
				),
				NewNode(2,
					NewNode(4, nil, nil),
					nil,
				),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InvertBinaryTree(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Test %q failed.\nExpected: %+v\nGot: %+v", tt.name, tt.expected, tt.input)
			}
		})
	}
}
