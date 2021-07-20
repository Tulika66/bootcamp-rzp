package main

import "fmt"

type Tree struct{
	valueAtIndex string
	leftNode *Tree
	rightNode *Tree
}

//preorder traversal for tree rooted at "tree"
//takes a pointer to the struct element "tree", visits current node,left and then right nodes recursively
func preorderTraversalTree(tree *Tree){
	if tree==nil{
		return
	}
    fmt.Println(tree.valueAtIndex)
	preorderTraversalTree(tree.leftNode)
	preorderTraversalTree(tree.rightNode)


}

//post order traversal for tree rooted at "tree"
//takes a pointer to the struct element "tree", visits current left , right and then current node recursively

func postorderTraversalTree(tree *Tree){
	if tree==nil{
		return
	}

	preorderTraversalTree(tree.leftNode)
	preorderTraversalTree(tree.rightNode)
	fmt.Println(tree.valueAtIndex)

}

func main() {
// example -: a + b
//build tree
root:=&Tree{"+",nil,nil}
root.leftNode = &Tree{"a",nil,nil}
root.rightNode=&Tree{"b",nil,nil}


preorderTraversalTree(root)

}
