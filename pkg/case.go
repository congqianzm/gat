package pkg

type Case struct {
	Name     string
	Path     string
	Steps    []Step
	Setup    []func()
	Teardown []func()
	Cache    Cache
}

func (c Case) Run() {
	for _, stepInstance := range c.Steps {
		stepInstance.Run(c)
	}
}
