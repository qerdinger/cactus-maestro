package cmastero_registry

import "fmt"

type PythonInterpreter struct {
	Language string "python3.12"
}

func (interpreter *PythonInterpreter) Build(code string) (string, error) {
	dockerfile := `
	FROM internal.docker.cmaestro.svc.cluster.local/python

	LABEL AUTHORS="Cactus Maestro (c) https://github.com/qerdinger/cactus-maestro - MIT License"
	LABEL IMG_LICENSE="MIT"

	RUN apt-get update
	`

	fmt.Println("Building python image...")
	fmt.Println(interpreter.Language)

	fmt.Println(dockerfile)
	return dockerfile, nil
}
