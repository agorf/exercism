package tree

import (
	"errors"
	"slices"
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
	if len(records) == 0 {
		return nil, nil
	}
	sortedRecords := append([]Record{}, records...)        // Copy
	slices.SortFunc(sortedRecords, func(a, b Record) int { // Sort in-place
		switch {
		case a.ID < b.ID:
			return -1
		case a.ID > b.ID:
			return 1
		default:
			switch {
			case a.Parent < b.Parent:
				return -1
			case a.Parent > b.Parent:
				return 1
			default:
				return 0
			}
		}
	})
	if sortedRecords[0].ID != 0 {
		return nil, errors.New("no root node")
	}
	if sortedRecords[0].Parent > 0 {
		return nil, errors.New("root node with parent")
	}
	nodes := map[int]*Node{}
	for i, record := range sortedRecords {
		if i != record.ID {
			return nil, errors.New("non-continuous")
		}
		if _, exists := nodes[record.ID]; exists {
			return nil, errors.New("duplicate node")
		}
		node := &Node{ID: record.ID}
		nodes[node.ID] = node
		if record.ID == 0 { // Root
			continue
		}
		if record.ID == record.Parent {
			return nil, errors.New("cycle-directly")
		}
		parentNode, exists := nodes[record.Parent]
		if !exists {
			return nil, errors.New("missing parent")
		}
		parentNode.Children = append(parentNode.Children, node)
	}
	return nodes[0], nil
}
