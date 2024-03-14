package pkg

import (
	"fmt"
	"strings"
)

type Request struct {
	Host   string
	Path   string
	Method string
	JSON   map[string]any
}

func (r Request) String() string {
	return fmt.Sprint(strings.Join([]string{r.Host, r.Path, r.Method, fmt.Sprint(r.JSON)}, "\n"))
}
