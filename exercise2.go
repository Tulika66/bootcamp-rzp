package main

import "fmt"

type Tree struct{
	valueAtIndex string
	leftNode *Tree
	rightNode *Tree
}

//Method of Tree struct for preorder traversal for tree rooted at "tree"
// visits first current node,left and then right node recursively
func (tree *Tree)preorderTraversalTree(){
	if tree==nil{
		return
	}
    fmt.Println(tree.valueAtIndex)

	if tree.leftNode != nil{
		tree.leftNode.preorderTraversalTree()
	}

	if tree.rightNode!=nil {
		tree.rightNode.preorderTraversalTree()
	}


}

//Method of Tree struct for post order traversal for tree rooted at "tree"
//visits first left , right and then current node recursively
func (tree *Tree)postorderTraversalTree(){
	if tree==nil{
		return
	}

	if tree.leftNode != nil{
		tree.leftNode.postorderTraversalTree()
	}

	if tree.rightNode!=nil {
		tree.rightNode.postorderTraversalTree()
	}

	fmt.Println(tree.valueAtIndex)

}


func main() {

	// example below -: a + b
	//build tree manually
	root:=&Tree{"+",nil,nil}
	root.leftNode = &Tree{"a",nil,nil}
	root.rightNode=&Tree{"b",nil,nil}

    fmt.Println("\nPreOrder Traversal :- ")
	root.preorderTraversalTree()

	fmt.Println("\nPostOrder Traversal :-")
    root.postorderTraversalTree()
}
