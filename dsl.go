package main

type TaskTable struct {
	name       string
	action     Action
	depends_on []*TaskTable
}

type Action func()

var NoAction Action = nil

func (this Action) Run() {
	if this != nil {
		this()
	}
}

type Deps map[string]*TaskTable

func (t *TaskTable) Run() {
	if t.depends_on != nil {
		for _, TaskTable := range t.depends_on {
			TaskTable.action()
		}
	}
	t.action.Run()
}

func (this *TaskTable) Deps(t ...*TaskTable) *TaskTable {
	for i := range t {
		this.depends_on = append(this.depends_on, t[i])
	}
	return this
}

func Task(name string, action func()) *TaskTable {
	TaskTable := &TaskTable{}
	TaskTable.name = name
	TaskTable.action = action
	return TaskTable
}

func main() {
	Action(func() {
		println("action A")
	}).Run()

	hello := Task("hello", func() {
		println("Hello")
	})

	world := Task("world", func() {
		println("World!")
	})

	greet := Task("greet", NoAction).Deps(hello, world)
	greet.Run()
}
