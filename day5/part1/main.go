package main

import (
  "strconv"
  "strings"
  "bufio"
  "fmt"
  "os"
  "io"
)


func step(stack []int, i int) int {
  stack[i] += 1
  fmt.Println(stack)
  return i + stack[i] - 1
}

func main() {
  file, _ := os.Open("../data")
  reader := bufio.NewReader(file)
  stack := make([]int, 0, 0)
  for {
    line, err := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    num, _ := strconv.Atoi(line)
    stack = append(stack, num)
    if err != nil {
      if err != io.EOF {
          println(" > Failed!: %v\n", err)
      }
      break
    }
  }
  steps := 0
  i := 0
  for {
    steps += 1
    i = step(stack, i)
    println(steps, i)
    if (i < 0 ) || (i > len(stack)-1) {
      break
    }
  }
}