package main

import "fmt"

func main() {
  s1 := []int{1,2,3,4}
  s2 := []int{9,8,7}
  // 当append之后，长度不超过源slice底层数组的长度，则会直接修改源slice中的数据
  // output: [1 9 8 7 9 8 7]
  fmt.Println(append(append(s1[:1], s2...), s1[1:]...)) //?

  s3 := []int{1,2,3}
  s4 := []int{9,8,7}
  // 当append之后，超过了slice底层数组的长度，则会分配一个新的更长的数组
  // output: [1 9 8 7 2 3]
  fmt.Println(append(append(s3[:1], s4...), s3[1:]...))
}
