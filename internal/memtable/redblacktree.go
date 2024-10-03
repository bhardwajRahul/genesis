package memtable

/*
Our memtable will use a Red-Black tree as its under-the-hood implementation
Meant to replace our original hash-table
*/

type color int

// Red = 0, Black = 1
const (
	RED color = iota
	BLACK
)

type Node struct {
	Key    string
	Value  string
	Parent *Node
	Left   *Node
	Right  *Node
	Color  color
}

type RedBlackTree struct {
	root *Node
}

func (tree *RedBlackTree) Insert(key, value string) {
	node := &Node{Key: key, Value: value, Color: RED}

	if tree.root == nil { // If tree is empty
		tree.root = node
	} else {
		currentNode := tree.root // Start from the root
		var parentNode *Node     // Contrast to standard BST, we need to keep track of parent nodes
		for currentNode != nil {
			parentNode = currentNode
			// Standard BST insertion
			if node.Key < currentNode.Key {
				currentNode = currentNode.Left
			} else {
				currentNode = currentNode.Right
			}
		}
		// now we're at a nil node, and parentNode is set to the last non-nil node we traversed, while currentNode is at nil
		node.Parent = parentNode
		// now we have to attach the node to our tree, again standard BST insertion
		if node.Key < parentNode.Key {
			parentNode.Left = node
		} else {
			parentNode.Right = node
		}
	}
	// Since this insertion may have violated RBT properties, we need to fix it
	tree.fixInsert(node)
}

func (tree *RedBlackTree) fixInsert(node *Node) {
	parentNode := node.Parent

	// While the parent node is red (RBT violation since newly added nodes are red by default),
	// we need to rotate and/or recolor
	for parentNode != nil && parentNode.Color == RED {
		grandParentNode := parentNode.Parent

		if parentNode == grandParentNode.Left {
			uncleNode := grandParentNode.Right

			if uncleNode.Color == RED {
				parentNode.Color = BLACK
				uncleNode.Color = BLACK
				grandParentNode.Color = RED
				node = grandParentNode
			} else { // uncle node is black
				if node == parentNode.Left {
					// node-parent-grandparent form a line, thus recolor & rotate grandparent right (opp. of node)
					parentNode.Color = BLACK
					grandParentNode.Color = RED
					rotateRight(grandParentNode)
				} else { // node is right child of parent node
					// node-parent-grandparent form a triangle, thus rotate parent left (opp. of node)
					node = parentNode
					rotateLeft(parentNode)
				}
			}
		} else { // Parent is right child of grandparent
			uncleNode := grandParentNode.Left

			if uncleNode.Color == RED {
				parentNode.Color = BLACK
				uncleNode.Color = BLACK
				grandParentNode.Color = RED
				node = grandParentNode
			} else {
				if node == parentNode.Left {
					// node-parent-grandparent form a line, thus recolor & rotate grandparent right (opp. of node)
					parentNode.Color = BLACK
					grandParentNode.Color = RED
					rotateRight(grandParentNode)
				} else { // node is right child of parent node
					// node-parent-grandparent form a triangle, thus rotate parent left (opp. of node)
					node = parentNode
					rotateLeft(parentNode)
				}
			}
		}
		parentNode = node.Parent // move node up since the violations will just cascade upwards
	}
	// Root of the tree must always be black
	tree.root.Color = BLACK
}

func rotateRight(node *Node) *Node {
	leftChild := node.Left      // store node's leftChild
	node.Left = leftChild.Right // overwrite node.Left with leftChild's right node
	if leftChild.Right != nil {
		leftChild.Right.Parent = node // reassign the parent to node
	}

	leftChild.Parent = node.Parent // move leftChild's parent up a level (bc its new position)
	if node.Parent == nil {
		// node is the root
		node = leftChild
	} else if node == node.Parent.Right {
		node.Parent.Right = leftChild
	} else {
		node.Parent.Left = leftChild
	}
	leftChild.Right = node  // move node down to be left child's right node
	node.Parent = leftChild // handle left child's left node

	return leftChild
}

func rotateLeft(node *Node) *Node {
	rightChild := node.Right
	node.Right = rightChild.Right
	if rightChild.Left != nil {
		rightChild.Left.Parent = node
	}
	rightChild.Parent = node.Parent
	if node.Parent == nil {
		node = rightChild
	} else if node == node.Parent.Left {
		node.Parent.Left = rightChild
	} else {
		node.Parent.Right = rightChild
	}
	rightChild.Left = node
	node.Parent = rightChild

	return rightChild
}

func (tree *RedBlackTree) Find(key string) {
	// impl
}
