package main

type IRobot interface {
	Run()
	Start()
	Stop()
}

type MyRobot struct {
	name string
}

func (r *MyRobot) Run() {
	println("My Robot Run")
}

func (r *MyRobot) Start() {
	println("My Robot Start")
}

func (r *MyRobot) Stop() {
	println("My Robot Stop")
}

func StartupRobot(robot IRobot) {
	// 断言robot是*MyRobot类型
	r, ok := robot.(*MyRobot)
	if ok {
		println("我叫" + r.name)
	}
}

func StartupRobot2(robot interface{}) {
	// 断言robot实现了IRobot接口
	r, ok := robot.(IRobot)
	if ok {
		r.Run()
	}
}

func main() {
	//var robot IRobot
	robot := &MyRobot{name: "萝卜头"}
	StartupRobot(robot)
	StartupRobot2(robot)
}
