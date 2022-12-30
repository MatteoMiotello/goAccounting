package bootstrap

type Initializer interface {
	Init() error
}
