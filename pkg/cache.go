package pkg

type Cache map[string]any

func (c Cache) Get(key string) any {
	return c[key]
}

func (c Cache) Set(key string, value any) {
	c[key] = value
}

func (c Cache) Del(key string) {
	delete(c, key)
}
