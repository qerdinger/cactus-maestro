package cmaestro_ingest

type Function struct {
	Name       string
	Args       []Argument
	ReturnType Primitive

	interpreter  string
	requirements []string
}

func NewFunction(name string, returnType string) *Function {
	f := Function{Name: name, Args: []Argument{}, ReturnType: Primitive(returnType)}

	// Interpreter
	// to be computed dynamically later
	f.interpreter = "python3.12"
	f.requirements = []string{}

	return &f
}

func (f *Function) AddArgument(name string, returnType string) *Function {
	f.Args = append(f.Args, Argument{Name: name, Type: Primitive(returnType)})
	return f
}
