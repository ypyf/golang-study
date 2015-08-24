package main

import (
    "fmt"
    //"os"
)

type Error interface {
    Throw()
}

type Program interface {
    Run()
    Init()
    Shutdown()
    Error
}

type Game struct {
}

func (g *Game) Run() {
    fmt.Println("游戏开始运行")
}

func (g *Game) Init() {
    fmt.Println("游戏初始化")
}

func (g *Game) Shutdown() {
    fmt.Println("游戏关闭")
}

func (g *Game) Throw() {
    fmt.Println("游戏出错")
}

func (g Game) String() string {
    return fmt.Sprintf("%s", "游戏程序")
}

func main() {
    var c Program
    
    c = &Game{}
    fmt.Println(c)
    c.Init()
    c.Run()
    c.Throw()
    c.Shutdown()
    
}

