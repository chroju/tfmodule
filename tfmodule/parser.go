package tfmodule

// Parser represents a terraform module parser.
type Parser interface {
	Parse(source string) (*Module, error)
}
