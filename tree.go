package tonho

// Location represents a location in the source code.
type Location interface {
	Start() int
	End() int

	// Text gets the text of the file at the location.
	Text() string

	// File gets the file name of the location.
	File() string
}

type Tree interface {
	// Location gets the location of the tree.
	Location() Location
}

// Node is a struct that represents a node in the concrete syntax tree.
type Node struct {
	Kind     int
	Children []Tree

	// The location of the node.
	location Location
}

// The kinds of nodes.
const (
	FileNode = iota
	ValNode
	VarNode
	WhileNode
	ForNode
	LoopNode
	ExprNode
	AssignNode
	FunNode
	StructNode
	EnumNode
	WhenNode
	IfNode
	ElseNode
	CallNode
	NumberNode
	StringNode
	BoolNode
	IdentifierNode
	ParameterNode
	TypeNameNode
	TypeApplicationNode
	GenericsNode
)

// Location gets the location of the node.
func (n Node) Location() Location {
	return n.location
}

// NewNode creates a new node with the given kind and children.
func NewNode(kind int, children []Tree) Node {
	return Node{Kind: kind, Children: children}
}
