package main

func loop(num int) int {
  current := 2
  layer := 2
  x := 1
  y := 0
  for {
    for i := 0; i < layer - 1; i++ {
      current += 1  
      if layer % 2 == 0 {
        y += 1
      } else {
        y -= 1
      }
      println(current, layer, x, y)
      if current == num {
        return steps(x, y)      
      }
    }
    for i := 0; i < layer; i++ {
      current += 1
      if layer % 2 == 0 {
        x -= 1
      } else {
        x += 1
      }
      println(current, layer,  x, y)
      if current == num {
        return steps(x, y)      
      }
    }
    layer += 1
  }
}

func abs(x int) int {
  if x > 0 {
    return x
  } else {
    return -x
  }
}

func steps(x, y int) int {
  return abs(x) + abs(y)
}

func main() {
  println(loop(361527))
}