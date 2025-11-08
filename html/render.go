package html

import (
	"net/http"
	"strings"
)

// String returns the HTML string representation of a node.
func String(n Node) string {
	var sb strings.Builder
	if err := n.Render(&sb); err != nil {
		return ""
	}
	return sb.String()
}

// Render writes the HTML to an http.ResponseWriter.
func Render(n Node, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return n.Render(w)
}
