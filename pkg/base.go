package pkg

type RunnerType string

const (
	Text RunnerType = "text"
	JSON RunnerType = "json"
	HTML RunnerType = "html"
)

type ActionType string

const (
	HTTP ActionType = "case"
)

type StatusType int

const (
	Fail    StatusType = 0
	Success StatusType = 1
	Running StatusType = 2
)
