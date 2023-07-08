package calc

type Calc struct {
	A, B   float64
	action IStrategyAction
}

func (c *Calc) SetStrategyAction(action IStrategyAction) {
	c.action = action
}

func (c Calc) DoAction() float64 {
	return c.action.action(c.A, c.B)
}

type IStrategyAction interface {
	action(a float64, b float64) float64
}

type Divade struct {
}

func (add Divade) action(a float64, b float64) float64 {
	if b == 0 {
		panic("Division by 0 is deprecated")
	}
	return a / b
}

type Substruct struct {
}

func (s Substruct) action(a float64, b float64) float64 {
	return a - b
}

type Add struct {
}

func (s Add) action(a float64, b float64) float64 {
	return a + b
}

type Multiply struct {
}

func (s Multiply) action(a float64, b float64) float64 {
	return a * b
}
