package pov

type Tree struct {
	value    string
	children []*Tree
}

func New(value string, children ...*Tree) *Tree {
	return &Tree{value, children}
}

func (t *Tree) Value() string {
	return t.value
}

func (t *Tree) Children() []*Tree {
	return t.children
}

func (t *Tree) String() string {
	if t == nil {
		return "nil"
	}
	return t.value + "\n" + t.buildBranches("", true)
}

func (t *Tree) buildBranches(prefix string, isLast bool) string {
	if len(t.children) == 0 {
		return ""
	}

	result := ""
	for i, child := range t.children {
		isLastChild := i == len(t.children)-1

		if isLastChild {
			result += prefix + "└── " + child.value + "\n"
		} else {
			result += prefix + "├── " + child.value + "\n"
		}

		var newPrefix string
		if isLastChild {
			newPrefix = prefix + "    "
		} else {
			newPrefix = prefix + "│   "
		}

		result += child.buildBranches(newPrefix, isLastChild)
	}

	return result
}

func (t *Tree) tracePath(target string, trace []*Tree) []*Tree {
	trace = append(trace, t)

	if t.value == target {
		return trace
	}

	for _, kid := range t.children {
		result := kid.tracePath(target, trace)
		if result != nil {
			return result
		}
	}

	return nil
}

func (t *Tree) FromPov(targetValue string) *Tree {
	trace := t.tracePath(targetValue, make([]*Tree, 0, 4))

	if len(trace) == 0 {
		return nil
	}

	for i := len(trace) - 1; i > 0; i-- {
		curr, parent := trace[i], trace[i-1]
		for j, kid := range parent.children {
			if kid == curr {
				parent.children = append(parent.children[:j], parent.children[j+1:]...)
				break
			}
		}

		curr.children = append(curr.children, parent)
	}

	return trace[len(trace)-1]
}

func (t *Tree) PathTo(from, to string) []string {
	fromTree := t.FromPov(from)

	if fromTree == nil {
		return nil
	}

	trace := fromTree.tracePath(to, make([]*Tree, 0, 4))
	result := make([]string, 0, len(trace))
	for _, tree := range trace {
		result = append(result, tree.value)
	}

	return result
}
