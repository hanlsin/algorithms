package main

import (
	"fmt"
)

type TreeNode struct {
	name *string
	children []*TreeNode
	childSize int
}

type Tree struct {
	categories map[string]*TreeNode
	root *TreeNode
}

type TreeError interface {
	Error() string
}

type InvalidArgumentError struct{}
func (e *InvalidArgumentError) Error() string {
	return "Invalid Argument"
}

func NewTreeNode(name *string) *TreeNode {
	return &TreeNode{
		name: name,
		children: make([]*TreeNode, 2),
		childSize: 0,
	}
}

func NewCategoryTree() *Tree {
	root := NewTreeNode(nil)
	return &Tree{
		categories: make(map[string]*TreeNode),
		root: root,
	}
}

func (t *Tree) AddCategory(parent *TreeNode, name string) (*TreeNode, error) {
	if parent == nil {
		return nil, &InvalidArgumentError{}
	}
	if parent.name != nil {
		if _, ok := t.categories[*parent.name]; !ok {
			return nil, &InvalidArgumentError{}
		}
	} else {
		if parent != t.root {
			return nil, &InvalidArgumentError{}
		}
	}
	if _, ok := t.categories[name]; ok {
		return nil, &InvalidArgumentError{}
	}

	newChild := NewTreeNode(&name)
	t.categories[name] = newChild

	if len(parent.children) == parent.childSize {
		nl := make([]*TreeNode, 2*parent.childSize)
		copy(nl, parent.children)
		parent.children = nl
	}
	parent.children[parent.childSize] = newChild
	parent.childSize++

	return newChild, nil
}

func getChildren (node *TreeNode) []*TreeNode {
	return node.children
}

func main() {
	fmt.Println("Hello World")

	tree := NewCategoryTree()
	a, _ := tree.AddCategory(tree.root, "A")
	tree.AddCategory(a, "B")
	c, _ := tree.AddCategory(a, "C")
	tree.AddCategory(c, "D")
	tree.AddCategory(c, "E")
	tree.AddCategory(c, "F")

	_, error := tree.AddCategory(c, "A")
	if error != nil {
		fmt.Println(error)
	}
	_, error = tree.AddCategory(NewTreeNode(nil), "NO EXIST")
	if error != nil {
		fmt.Println(error)
	}
	name := "no-exist"
	_, error = tree.AddCategory(NewTreeNode(&name), "NO EXIST")
	if error != nil {
		fmt.Println(error)
	}

	for i, ch := range getChildren(a) {
		fmt.Printf("[%02d] child: %v\n", i, *ch.name)
	}

	for i, ch := range getChildren(c) {
		if ch == nil {
			break
		}
		fmt.Printf("[%02d] child: %v\n", i, *ch.name)
	}
}
