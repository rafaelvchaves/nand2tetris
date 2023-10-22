package token

type A struct {
	token
	Value string
}

type C struct {
	token
	Dest string
	Comp string
	Jump string
}

type Label struct {
	token
	Name string
}

type Empty struct {
	token
}

type Token interface {
	tokenSumType()
}

type token interface {
	Token
}
