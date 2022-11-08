package strategies

import "strategy/core"

type BallroomDancing struct {
}

func (b BallroomDancing) Name() string {
	return `Ballroom dancing`
}

func (b BallroomDancing) TestChild(child core.Child) bool {
	return child.Predisposition() == `body`
}
