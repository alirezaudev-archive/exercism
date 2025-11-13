package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// Handle empty input
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))

	for i, record := range records {
		if record.ID != i {
			return nil, errors.New("non-continuous")
		}

		if nodes[record.ID] != nil {
			return nil, errors.New("duplicate node")
		}

		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, errors.New("root node has parent")
			}
			nodes[0] = &Node{ID: 0}
			continue
		}

		if record.Parent >= record.ID {
			return nil, errors.New("invalid parent")
		}

		parent := nodes[record.Parent]
		if parent == nil {
			return nil, errors.New("parent not found")
		}

		child := &Node{ID: record.ID}
		nodes[record.ID] = child
		parent.Children = append(parent.Children, child)
	}

	return nodes[0], nil
}
