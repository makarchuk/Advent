package main

func main() {
  loop()
}

func loop() int {
  var a [50][50]int
  x := 25
  y := 25
  num := 361527
  layer := 2
  current := 0
  a[x][y] = 1
  a[26][25] = 1
  x = 26
  for {
    for i := 0; i < layer - 1; i++ {
        if layer % 2 == 0 {
          y += 1
        } else {
          y -= 1
        }
        current = sum(a, x, y)
        a[x][y] = current
        println(current, layer, x, y)
        if current > num {
          return current
        }
      }
      for i := 0; i < layer; i++ {
        if layer % 2 == 0 {
          x -= 1
        } else {
          x += 1
        }
        current = sum(a, x, y)
        a[x][y] = current
        println(current, layer,  x, y)
        if current > num {
          return current
        }
      }
      layer += 1
    }    
}

func sum(grid [50][50]int, x int, y int) int {
  up := grid[x+1][y+1] + grid[x][y+1] + grid[x-1][y+1]
  mid := grid[x+1][y] + grid[x-1][y]
  down := grid[x+1][y-1] + grid[x][y-1] + grid[x-1][y-1]
  println("SUM", up, mid, down)
  return up + down + mid
}