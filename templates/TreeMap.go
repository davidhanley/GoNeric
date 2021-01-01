package main


type TreeMapType1 int
type TreeMapType2 int

type Node struct {
	key    TreeMapType1
	val    TreeMapType2
	left   *Node
	right  *Node
	height int
}

func copy(node *Node) *Node {
	if node == nil {
		return nil
	}
	n := new(Node) //should be a bitwise copy, this returns a zeroed out node, then copies over it
	*n = *node
	return n
	//return &Node{key: node.key, right: node.right, left: node.left, height: node.height, val: node.val}
}

func insert(node *Node, key TreeMapType1, val TreeMapType2) *Node {
	if node == nil {
		return &Node{key: key, val: val}
	}
	if node.key > key {
		node = copy(node)
		node.left = insert(node.left, key, val)
	} else if node.key < key {
		node = copy(node)
		node.right = insert(node.right, key, val)
	} else {
		return &Node{key: key}
	}
	return rebalance(node)
}

func delete(node *Node, key TreeMapType1) *Node {
	if node == nil {
		return node
	}

	if node.key > key {
		node = copy(node)
		node.left = delete(node.left, key)
	} else if node.key < key {
		node = copy(node)
		node.right = delete(node.right, key)
	} else {
		if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			node = copy(node)
			mostLeftChild := findMostLeftChild(node.right)
			node.key = mostLeftChild.key
			node.right = delete(node.right, node.key)
		}
	}

	if node != nil {
		node = rebalance(node)
	}
	return node
}

func findMostLeftChild(current *Node) *Node {
	for ; current.left != nil; current = current.left {
	}
	return current
}

//z is already a copy, so we can mute it
func rebalance(z *Node) *Node {
	updateHeight(z)
	balance := getBalance(z)
	if balance > 1 {
		if height(z.right.right) > height(z.right.left) {
			z = rotateLeft(z)
		} else {
			z.right = rotateRight(copy(z.right))
			z = rotateLeft(z)
		}
	} else if balance < -1 {
		if height(z.left.left) > height(z.left.right) {
			z = rotateRight(z)
		} else {
			z.left = rotateLeft(copy(z.left))
			z = rotateRight(z)
		}
	}
	return z
}

func rotateRight(y *Node) *Node {
	x := copy(y.left)
	z := x.right // copy needed?
	x.right = y
	y.left = z
	updateHeight(y)
	updateHeight(x)
	return x
}

func rotateLeft(y *Node) *Node {
	x := copy(y.right)
	z := x.left // copy needed?
	x.left = y
	y.right = z
	updateHeight(y)
	updateHeight(x)
	return x
}

func updateHeight(n *Node) {
	n.height = 1 + int(max(minmaxType1(height(n.left)), minmaxType1(height(n.right))))
}

func height(n *Node) int {
	if n == nil {
		return -1
	} else {
		return n.height
	}
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	} else {
		return height(n.right) - height(n.left)
	}
}

func toStream(n *Node, c chan TreeMapType1) {
	if n != nil {
		toStream(n.left, c)
		c <- n.key
		toStream(n.right, c)
	}
}

func size(n *Node) int {
	if n == nil {
		return 0
	}

	return 1 + size(n.left) + size(n.right)
}
