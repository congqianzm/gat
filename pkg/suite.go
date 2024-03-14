package pkg

type Suite struct {
	Name     string
	Path     string
	Cases    []Case
	Setup    []func()
	Teardown []func()
	Cache    Cache
}

func (s Suite) Run() {
	for _, caseInstance := range s.Cases {
		caseInstance.Run()
	}
}
