package kindergarten

import (
	"fmt"
	"math/rand"
	"strategy/core"
	"time"
)

type Child struct {
	age            int
	predisposition string
	sport          core.Sport
}

func (c Child) Age() int {
	return c.age
}

func (c *Child) Learn(s core.Sport) {
	c.sport = s
	fmt.Printf(" --- Child with age %d have learnd %s\n", c.Age(), s.Name())
}

func (c Child) Predisposition() string {
	return c.predisposition
}

func GenerateChildren() Child {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(6)
	p := PredispositionGenerator()

	return Child{age: a, predisposition: p}
}

func PredispositionGenerator() string {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(3)

	switch r {
	case 0:
		return `strength`
	case 1:
		return `body`
	case 2:
		return `brain`
	default:
		fallthrough
	case 3:
		return `nothing`
	}
}
