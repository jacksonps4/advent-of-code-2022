package main

type crt struct {
	Bits     []bool
	position int
}

func NewCRT() *crt {
	return &crt{
		Bits: make([]bool, 240),
	}
}

func (c *crt) Cycle() {
	c.position++
}

func (c *crt) Position() int {
	return c.position
}

func (c *crt) Set() {
	c.Bits[c.position] = true
}

func (c *crt) Unset() {
	c.Bits[c.position] = false
}

func (c *crt) String() string {
	data := ""
	for i, c := range c.Bits {
		if c {
			data = data + "#"
		} else {
			data = data + "."
		}
		if i+1 > 0 && (i+1)%40 == 0 {
			data = data + "\n"
		}
	}
	return data
}
