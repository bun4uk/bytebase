// Package ast defines the abstract syntax tree of mybatis mapper xml.
package ast

import (
	"encoding/xml"
	"io"
	"strings"
)

var (
	_ Node = (*QueryNode)(nil)
	_ Node = (*TextNode)(nil)
)

// QueryNodeType is the type of the query node.
type QueryNodeType uint

const (
	// QueryNodeTypeSelect represents a select query node.
	QueryNodeTypeSelect QueryNodeType = iota
	// QueryNodeTypeInsert represents a insert query node.
	QueryNodeTypeInsert
	// QueryNodeTypeUpdate represents a update query node.
	QueryNodeTypeUpdate
	// QueryNodeTypeDelete represents a delete query node.
	QueryNodeTypeDelete
)

// QueryNode represents a query node.
type QueryNode struct {
	// ID is the id of the query node.
	ID string
	// Type is the type of the query node.
	Type QueryNodeType
	// Children is the children of the query node.
	Children []Node
}

// RestoreSQL implements Node interface.
func (n *QueryNode) RestoreSQL(ctx *RestoreContext, w io.Writer) error {
	var sb strings.Builder

	for _, node := range n.Children {
		if err := node.RestoreSQL(ctx, &sb); err != nil {
			return err
		}
	}
	stmt := sb.String()
	trimmed := strings.TrimSpace(stmt)
	if len(trimmed) == 0 {
		return nil
	}
	if _, err := w.Write([]byte(trimmed)); err != nil {
		return err
	}
	if _, err := w.Write([]byte(";\n")); err != nil {
		return err
	}
	return nil
}

// AddChild adds a child to the query node.
func (n *QueryNode) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

// NewQueryNode creates a new query node.
func NewQueryNode(startEle *xml.StartElement) *QueryNode {
	n := &QueryNode{}
	switch startEle.Name.Local {
	case "select":
		n.Type = QueryNodeTypeSelect
	case "insert":
		n.Type = QueryNodeTypeInsert
	case "update":
		n.Type = QueryNodeTypeUpdate
	case "delete":
		n.Type = QueryNodeTypeDelete
	}

	for _, attr := range startEle.Attr {
		if attr.Name.Local == "id" {
			n.ID = attr.Value
		}
	}
	return n
}
