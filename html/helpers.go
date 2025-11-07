package html

// If conditionally renders children when condition is true.
func If(condition bool, children ...any) Node {
	if condition {
		return Group(children...)
	}
	return Group()
}

// IfExec conditionally executes a function and renders result when condition is true.
func IfExec(condition bool, fn func() Node) Node {
	if condition {
		return fn()
	}
	return Group()
}

// IfElse renders trueChildren when condition is true, otherwise falseChildren.
func IfElse(condition bool, trueChildren Node, falseChildren Node) Node {
	if condition {
		return trueChildren
	}
	return falseChildren
}

// IfElseExec executes trueFn when condition is true, otherwise falseFn.
func IfElseExec(condition bool, trueFn func() Node, falseFn func() Node) Node {
	if condition {
		return trueFn()
	}
	return falseFn()
}

// IfNot conditionally renders children when condition is false.
func IfNot(condition bool, children ...any) Node {
	if !condition {
		return Group(children...)
	}
	return Group()
}

// Map transforms a slice into nodes using the provided function.
func Map[T any](items []T, fn func(T) Node) Node {
	nodes := make([]any, len(items))
	for i, item := range items {
		nodes[i] = fn(item)
	}
	return Group(nodes...)
}
