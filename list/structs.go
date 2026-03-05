package list

type Node struct {
	Value  int
	Next   *Node
	Prev   *Node
	Random *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}
