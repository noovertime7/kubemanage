package runtime

var CloserHandler = NewCloserHandler()

type Closer interface {
	Close() error
	HandlerErr(error)
}

type closer struct {
	closers []Closer
}

func NewCloserHandler() *closer {
	return &closer{closers: make([]Closer, 0)}
}

func (c *closer) AddCloser(cls Closer) {
	c.closers = append(c.closers, cls)
}

func (c *closer) Close() {
	for _, cls := range c.closers {
		if err := cls.Close(); err != nil {
			cls.HandlerErr(err)
		}
	}
}
