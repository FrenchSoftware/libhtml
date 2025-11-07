[![Discord](https://img.shields.io/badge/Discord-Join%20Us-7289DA?logo=discord&logoColor=white)](https://discord.gg/MaqTgPF3Ch)

# libhtml

> A simple Go library to write HTML in plain Go

## Installation

```bash
go get github.com/frenchsoftware/libhtml
```

## Example

```go
package main

import (
	"fmt"
	"github.com/frenchsoftware/libhtml/attr"
	"github.com/frenchsoftware/libhtml/html"
)

func main() {
	page := html.Html(
		attr.Lang("en"),
		html.Head(
			html.Title(html.Text("Hello, World!")),
			html.Meta(attr.Charset("utf-8")),
		),
		html.Body(
			html.H1(html.Text("Welcome")),
			html.P(html.Text("This is a simple example.")),
		),
	)

	fmt.Println(html.String(page))
}
```

**Output:**
```html
<html lang="en"><head><title>Hello, World!</title><meta charset="utf-8"></head><body><h1>Welcome</h1><p>This is a simple example.</p></body></html>
```

## License

[MIT License](./LICENSE)
