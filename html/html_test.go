package html_test

import (
	"strings"
	"testing"

	"github.com/frenchsoftware/libhtml/attr"
	"github.com/frenchsoftware/libhtml/html"
)

func TestCompleteHTMLRendering(t *testing.T) {
	users := []string{"Alice", "Bob", "Charlie"}
	isLoggedIn := true

	doc := html.Html(
		attr.Lang("en"),
		html.Head(
			html.Title(html.Text("Users")),
			html.Meta(attr.Charset("utf-8")),
		),
		html.Body(
			html.H1(html.Text("User List")),
			html.If(isLoggedIn,
				html.Div(
					attr.Class("users"),
					html.Ul(
						html.Map(users, func(user string) html.Node {
							return html.Li(html.Text(user))
						}),
					),
				),
			),
			html.IfElse(len(users) > 0,
				html.P(html.Text("Users found")),
				html.P(html.Text("No users")),
			),
		),
	)

	var buf strings.Builder
	if err := doc.Output(&buf); err != nil {
		t.Fatalf("Output failed: %v", err)
	}

	htmlStr := buf.String()

	// Assert structure
	if !strings.Contains(htmlStr, `<html lang="en">`) {
		t.Errorf("Expected html tag with lang attribute, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `<meta charset="utf-8">`) {
		t.Errorf("Expected meta tag with charset, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `<li>Alice</li>`) {
		t.Errorf("Expected list item with Alice, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `<li>Bob</li>`) {
		t.Errorf("Expected list item with Bob, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `<li>Charlie</li>`) {
		t.Errorf("Expected list item with Charlie, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `Users found`) {
		t.Errorf("Expected 'Users found' text, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, `<div class="users">`) {
		t.Errorf("Expected div with users class, got: %s", htmlStr)
	}
}

func TestHelperFunctions(t *testing.T) {
	// Test IfNot
	notLoggedIn := false
	result := html.IfNot(notLoggedIn, html.P(html.Text("You are logged in")))
	htmlStr := html.String(result)
	if !strings.Contains(htmlStr, "You are logged in") {
		t.Errorf("IfNot failed: %s", htmlStr)
	}

	// Test IfExec
	count := 5
	result = html.IfExec(count > 0, func() html.Node {
		return html.P(html.Text("Count is positive"))
	})
	htmlStr = html.String(result)
	if !strings.Contains(htmlStr, "Count is positive") {
		t.Errorf("IfExec failed: %s", htmlStr)
	}

	// Test IfElseExec
	result = html.IfElseExec(count > 10,
		func() html.Node { return html.P(html.Text("Large")) },
		func() html.Node { return html.P(html.Text("Small")) },
	)
	htmlStr = html.String(result)
	if !strings.Contains(htmlStr, "Small") {
		t.Errorf("IfElseExec failed: %s", htmlStr)
	}

	// Test Group
	group := html.Group(
		html.P(html.Text("First")),
		html.P(html.Text("Second")),
		html.P(html.Text("Third")),
	)
	htmlStr = html.String(group)
	if !strings.Contains(htmlStr, "First") || !strings.Contains(htmlStr, "Second") || !strings.Contains(htmlStr, "Third") {
		t.Errorf("Group failed: %s", htmlStr)
	}
}

func TestStringConvenience(t *testing.T) {
	doc := html.Div(
		attr.Class("container"),
		html.P(html.Text("Hello, World!")),
	)

	htmlStr := html.String(doc)

	if !strings.Contains(htmlStr, `<div class="container">`) {
		t.Errorf("String() failed: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, "Hello, World!") {
		t.Errorf("String() failed: %s", htmlStr)
	}
}

func TestDocument(t *testing.T) {
	doc := html.Document(
		html.Html(
			attr.Lang("en"),
			html.Head(
				html.Title(html.Text("Test Page")),
				html.Meta(attr.Charset("utf-8")),
			),
			html.Body(
				html.H1(html.Text("Hello, World!")),
			),
		),
	)

	htmlStr := html.String(doc)

	// Check for DOCTYPE
	if !strings.HasPrefix(htmlStr, "<!DOCTYPE html>") {
		t.Errorf("Expected DOCTYPE declaration at the start, got: %s", htmlStr)
	}

	// Check for html tag with lang attribute
	if !strings.Contains(htmlStr, `<html lang="en">`) {
		t.Errorf("Expected html tag with lang attribute, got: %s", htmlStr)
	}

	// Check for head and body elements
	if !strings.Contains(htmlStr, "<head>") {
		t.Errorf("Expected head element, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, "<body>") {
		t.Errorf("Expected body element, got: %s", htmlStr)
	}

	// Check for content
	if !strings.Contains(htmlStr, "Hello, World!") {
		t.Errorf("Expected 'Hello, World!' text, got: %s", htmlStr)
	}
}

func TestDocumentEmpty(t *testing.T) {
	doc := html.Document()
	htmlStr := html.String(doc)

	// Check for DOCTYPE and empty html tag
	if !strings.HasPrefix(htmlStr, "<!DOCTYPE html>") {
		t.Errorf("Expected DOCTYPE declaration at the start, got: %s", htmlStr)
	}
	if !strings.Contains(htmlStr, "<html>") {
		t.Errorf("Expected html tag, got: %s", htmlStr)
	}
}
