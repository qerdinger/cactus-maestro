package cmastero_registry

type Factory interface {
	Interpreter() string
	Build(code string) (string, error)
}
