package main

import (
	"fmt"
)




func main() {
    // Manually building the following binary tree:
    //       1
    //      / \
    //     2   3
    //    /
    //   4

    tree := &BinaryTreeNode{
        Value: 1,
        LeftChild: &BinaryTreeNode{
            Value: 2,
            LeftChild: &BinaryTreeNode{Value: 4},
            RightChild: &BinaryTreeNode{Value: 5},
        },
        RightChild: &BinaryTreeNode{
            Value: 3,
            LeftChild: &BinaryTreeNode{
				Value: 6,
				LeftChild: &BinaryTreeNode{Value: 8},
				RightChild: &BinaryTreeNode{Value: 9},
			},
            RightChild: &BinaryTreeNode{Value: 7},
        },
    }


    printBinaryTree(tree, "", false)
	fmt.Println("------------------------------------")
	InvertBinaryTree(tree)
	printBinaryTree(tree, "", false)

}



type BinaryTreeNode struct {
	Value 		int;
	LeftChild 	*BinaryTreeNode;
	RightChild 	*BinaryTreeNode;
}

func InvertBinaryTree (node *BinaryTreeNode) {
	if node == nil {
		return
	}

	temp := node.RightChild	
	node.RightChild = node.LeftChild
	node.LeftChild = temp
	
	InvertBinaryTree(node.LeftChild)
	InvertBinaryTree(node.RightChild)
}

func printBinaryTree(node *BinaryTreeNode, prefix string, isLeft bool) {
    if node == nil {
        return
    }

    if node.RightChild != nil {
        printBinaryTree(node.RightChild, prefix+"    ", false)
    }

    if isLeft {
        fmt.Printf("%s└── %d\n", prefix, node.Value)
    } else {
        fmt.Printf("%s┌── %d\n", prefix, node.Value)
    }

    if node.LeftChild != nil {
        printBinaryTree(node.LeftChild, prefix+"    ", true)
    }
}
