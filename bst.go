package main

import (
	"fmt"
)

type BinaryST interface {
	Insert(int)
	Remove() int
	Find() bool
	Depth() int
}

type Node struct {
	Data       int
	LeftChild  *Node
	RightChild *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{Root: nil}
}

func (tree *BST) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Data: value}
		return
	}
	tree.Root.insert(value)
}

func (node *Node) insert(value int) {
	if value < node.Data {
		if node.LeftChild == nil {
			node.LeftChild = &Node{Data: value}
		} else {
			node.LeftChild.insert(value)
		}

	} else if value > node.Data {
		if node.RightChild == nil {
			node.RightChild = &Node{Data: value}
		} else {
			node.RightChild.insert(value)
		}
	}
}

func (tree *BST) Find(value int) bool {
	if tree.Root == nil {
		return false
	}
	return tree.Root.find(value)
}

func (node *Node) find(value int) bool {
	if node == nil {
		return false
	}
	if value == node.Data {
		return true
	} else if value > node.Data {
		return node.RightChild.find(value)
	} else {
		return node.RightChild.find(value)
	}
}

func (tree *BST) Depth() int {
	if tree.Root == nil {
		return 0
	}
	return tree.Root.countLevels()
}

func (node *Node) countLevels() int {
	if node == nil {
		return 0
	}
	leftDepth := node.LeftChild.countLevels()
	rightDepth := node.RightChild.countLevels()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (tree *BST) Remove(value int) bool {
	if tree.Root == nil {
		return false
	}
	var removed bool
	tree.Root, removed = remove(tree.Root, value)
	return removed
}

func remove(node *Node, value int) (*Node, bool) {
	if node == nil {
		return nil, false
	}

	var removed bool
	if node.Data > value {
		node.LeftChild, removed = remove(node.LeftChild, value)
	} else if node.Data < value {
		node.RightChild, removed = remove(node.RightChild, value)
	} else {
		// когда нашли нужный узел
		// если нет одного из детей или никого:

		removed = true
		if node.LeftChild == nil {
			return node.RightChild, true
		} else if node.RightChild == nil {
			return node.LeftChild, true
		}

		// если у узла 2 ребенка
		// найти минимаоьное поддерево справа, тк после удаления элемента на его место встанет больший из детей

		minNode := findMin(node.RightChild)
		node.Data = minNode.Data
		node.RightChild, _ = remove(node.RightChild, minNode.Data)
	}
	return node, removed
}

func findMin(node *Node) *Node {
	current := node
	for current.LeftChild != nil {
		current = current.LeftChild
	}
	return current
}

func (bst *BST) InOrder() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.inOrder(&result)
	}
	return result
}

func (node *Node) inOrder(result *[]int) {
	if node == nil {
		return
	}
	node.LeftChild.inOrder(result)
	*result = append(*result, node.Data)
	node.RightChild.inOrder(result)
}

func (bst *BST) PreOrder() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.preOrder(&result)
	}
	return result
}

func (node *Node) preOrder(result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Data)
	node.LeftChild.preOrder(result)
	node.RightChild.preOrder(result)
}

func (bst *BST) PostOrder() []int {
	var result []int
	if bst.Root != nil {
		bst.Root.postOrder(&result)
	}
	return result
}

func (node *Node) postOrder(result *[]int) {
	if node == nil {
		return
	}
	node.LeftChild.postOrder(result)
	node.RightChild.postOrder(result)
	*result = append(*result, node.Data)
}

func (tree *BST) PrintTree() []string {
	if tree.Root == nil {
		return []string{"Пустое дерево"}
	}

	return tree.Root.printLevels("", false)
}

func (node *Node) printLevels(prefix string, hasParent bool) []string {
	if node == nil {
		return []string{"Пустое дерево"}
	}
	var result []string
	childPrefix := prefix

	if hasParent == false { // корень
		result = append(result, fmt.Sprintf("%s%d", prefix, node.Data))
		// result = append(result, fmt.Sprintf("%s", prefix))
	} else {

		result = append(result, fmt.Sprintf("%s%d", prefix, node.Data))

	}
	if hasParent == false {
		childPrefix += "\\___"
	} else {
		childPrefix += "|___"
	}

	if node.LeftChild != nil {
		leftLines := node.LeftChild.printLevels(childPrefix, true)
		result = append(result, leftLines...)
	}

	if node.RightChild != nil {
		rightLines := node.RightChild.printLevels(childPrefix, true)
		result = append(result, rightLines...)
	}

	return result
}

func output(bst *BST) {
	lines := bst.PrintTree()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {

	bst := NewBST()

	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, value := range values {
		bst.Insert(value)
	}
	fmt.Println("Исходное дерево:")
	output(bst)
	fmt.Println("Прямой обход:\n", bst.PreOrder())

	fmt.Println("Центрированный обход:\n", bst.InOrder())
	fmt.Println("Обратный обход:\n", bst.PostOrder())
	fmt.Println("Глубина:", bst.Depth())

	fmt.Println("\nУдаляем 2:", bst.Remove(2))
	fmt.Println("После удаления 2:", bst.InOrder())

	output(bst)

	fmt.Println("\nУдаляем 3:", bst.Remove(3))
	fmt.Println("После удаления 3:", bst.InOrder())

	output(bst)

	fmt.Println("\nУдаляем 5:", bst.Remove(5))
	fmt.Println("После удаления 5:", bst.InOrder())
	fmt.Println("Глубина:", bst.Depth())

	output(bst)

	fmt.Println("\nПоиск 6:", bst.Find(6))
	fmt.Println("Поиск 5:", bst.Find(5))
}
