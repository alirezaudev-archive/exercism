package react

type cellType int

const (
	Input cellType = iota
	Compute
	Normal
)

type cell struct {
	kind           cellType
	value          int
	compute1       func(int) int
	compute2       func(int, int) int
	dep1           Cell
	dep2           Cell
	dependents     []*cell
	callbacks      map[int]func(int)
	nextCallbackID int
	r              *reactor
}

type canceler struct {
	c  *cell
	id int
}

type reactor struct {
	cells []*cell
}

func (c *canceler) Cancel() {
	if c.c == nil {
		return
	}

	delete(c.c.callbacks, c.id)
	c.c = nil
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(v int) {
	if c.kind != Input {
		return
	}
	if c.value == v {
		return
	}
	c.value = v
	c.r.propagateFrom(c)
}

func (c *cell) AddCallback(cb func(int)) Canceler {
	if c.callbacks == nil {
		c.callbacks = make(map[int]func(int))
	}
	id := c.nextCallbackID
	c.nextCallbackID++

	c.callbacks[id] = cb

	return &canceler{
		c:  c,
		id: id,
	}
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	c := &cell{
		kind:  Input,
		value: initial,
		r:     r,
	}
	r.cells = append(r.cells, c)
	return c
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	c := &cell{
		kind:     Compute,
		compute1: compute,
		dep1:     dep,
		r:        r,
	}

	if d, ok := dep.(*cell); ok {
		d.dependents = append(d.dependents, c)
	}

	c.recompute()
	r.cells = append(r.cells, c)
	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	c := &cell{
		kind:     Compute,
		compute2: compute,
		dep1:     dep1,
		dep2:     dep2,
		r:        r,
	}
	if d1, ok := dep1.(*cell); ok {
		d1.dependents = append(d1.dependents, c)
	}
	if d2, ok := dep2.(*cell); ok {
		d2.dependents = append(d2.dependents, c)
	}
	c.recompute()
	r.cells = append(r.cells, c)
	return c
}

func (c *cell) recompute() bool {
	if c.kind != Compute {
		return false
	}
	old := c.value

	var newVal int
	if c.compute1 != nil {
		newVal = c.compute1(c.dep1.Value())
	} else {
		newVal = c.compute2(c.dep1.Value(), c.dep2.Value())
	}

	if newVal == old {
		return false
	}

	c.value = newVal
	return true
}

func (r *reactor) propagateFrom(start *cell) {
	queue := []*cell{start}
	visited := make(map[*cell]bool)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dep := range cur.dependents {
			if visited[dep] {
				continue
			}
			visited[dep] = true

			if dep.kind != Compute {
				continue
			}

			changed := dep.recompute()
			if changed {
				queue = append(queue, dep)
				for _, cb := range dep.callbacks {
					cb(dep.value)
				}
			}
		}
	}
}
