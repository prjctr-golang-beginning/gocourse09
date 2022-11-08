package core

type Sport interface {
	Name() string
	TestChild(Child) bool
}

type Child interface {
	Age() int
	Predisposition() string
	Learn(Sport)
}
