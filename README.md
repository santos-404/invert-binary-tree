# Invert binary tree

This repository implements a simple function to invert a binary tree.

## how to run

To run the program, use the following command:

```bash
go run .
```

The program can also be built using:

```bash
go build
```

## how to test

Run the test suite with:

```bash
go test -v
```

## implementation

Below is the implementation of the inversion function. It's just a simple recursive operation.

```go
func InvertBinaryTree(node *BinaryTreeNode) {
	if node == nil {
		return
	}

	temp := node.RightChild
	node.RightChild = node.LeftChild
	node.LeftChild = temp

	InvertBinaryTree(node.LeftChild)
	InvertBinaryTree(node.RightChild)
}
```

## visualization

Example of binary tree inversion:

```
        ┌── 7
    ┌── 2
            ┌── 9
        └── 6
            └── 8
└── 1
        ┌── 5
    └── 2
        └── 4
```

becomes:

```
        ┌── 4
    ┌── 2
        └── 5
└── 1
            ┌── 8
        ┌── 6
            └── 9
    └── 3
        └── 7
```

