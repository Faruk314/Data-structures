package list

type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}