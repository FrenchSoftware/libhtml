package html

import "io"

// Node represents an HTML element or text node.
type Node interface {
	Output(w io.Writer) error
}

// Attribute represents an HTML attribute.
type Attribute interface {
	Key() string
	Value() string
}

// element represents a concrete HTML element with a tag, attributes, and children.
type element struct {
	tag         string
	attributes  []Attribute
	children    []Node
	selfClosing bool
}

// attribute represents a concrete HTML attribute with a key-value pair.
type attribute struct {
	key   string
	value string
}

// textNode represents a text node in the HTML tree.
type textNode struct {
	content string
}

// rawNode represents a raw HTML node that outputs content without escaping.
type rawNode struct {
	content string
}

// groupNode represents a group of nodes without a wrapper element.
type groupNode struct {
	children []Node
}

// documentNode represents a complete HTML document with DOCTYPE.
type documentNode struct {
	htmlElement Node
}

// Text creates a text node with the given content.
func Text(content string) Node {
	return &textNode{content: content}
}

// Raw creates a raw HTML node that outputs content without escaping.
// WARNING: Only use this with trusted content to avoid XSS vulnerabilities.
func Raw(content string) Node {
	return &rawNode{content: content}
}

// Output writes the text content to the writer.
func (t *textNode) Output(w io.Writer) error {
	_, err := w.Write([]byte(t.content))
	return err
}

// Output writes the raw HTML content to the writer without escaping.
func (r *rawNode) Output(w io.Writer) error {
	_, err := w.Write([]byte(r.content))
	return err
}

// Output writes the group's children to the writer.
func (g *groupNode) Output(w io.Writer) error {
	for _, child := range g.children {
		if err := child.Output(w); err != nil {
			return err
		}
	}
	return nil
}

// Output writes the DOCTYPE declaration followed by the HTML element.
func (d *documentNode) Output(w io.Writer) error {
	if _, err := w.Write([]byte("<!DOCTYPE html>")); err != nil {
		return err
	}
	return d.htmlElement.Output(w)
}

// Key returns the attribute's key.
func (a *attribute) Key() string {
	return a.key
}

// Value returns the attribute's value.
func (a *attribute) Value() string {
	return a.value
}

// Output writes the element to the writer.
func (e *element) Output(w io.Writer) error {
	// Opening tag
	if _, err := w.Write([]byte("<" + e.tag)); err != nil {
		return err
	}

	// Attributes
	for _, attr := range e.attributes {
		if _, err := w.Write([]byte(" " + attr.Key() + "=\"" + attr.Value() + "\"")); err != nil {
			return err
		}
	}

	if e.selfClosing {
		_, err := w.Write([]byte(">"))
		return err
	}

	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}

	// Children
	for _, child := range e.children {
		if err := child.Output(w); err != nil {
			return err
		}
	}

	// Closing tag
	_, err := w.Write([]byte("</" + e.tag + ">"))
	return err
}

// Element creates a new element with the given tag and children/attributes.
func Element(tag string, selfClosing bool, items ...any) Node {
	var attrs []Attribute
	var children []Node

	for _, item := range items {
		switch v := item.(type) {
		case Attribute:
			attrs = append(attrs, v)
		case Node:
			children = append(children, v)
		}
	}

	return &element{
		tag:         tag,
		attributes:  attrs,
		children:    children,
		selfClosing: selfClosing,
	}
}

// Group groups multiple nodes into a single node without a wrapper element.
func Group(items ...any) Node {
	var children []Node

	for _, item := range items {
		if node, ok := item.(Node); ok {
			children = append(children, node)
		}
	}

	return &groupNode{children: children}
}

// Attr creates a new attribute with the given key and value.
func Attr(key, value string) Attribute {
	return &attribute{key: key, value: value}
}
