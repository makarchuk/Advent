package main
import (
  "strings"
  "bufio"
  "sort"
  "fmt"
  "os"
  "io"
)

func validate(passes []string) int {
  counter := 0
  for _, passphrase := range(passes) {
    if validate_single(passphrase) {
      counter += 1
    } else {
      // println(passphrase)
    }
  }
  return counter
}

func validate_single(pass string) bool {
  uniqs := make(map[string]int)
  chunks := strings.Split(pass, " ")
  for _, val := range chunks {
    pieces := strings.Split(val, "")
    sort.Strings(pieces)
    key := strings.Join(pieces, "")
    uniqs[key] = 0
  }
  return len(uniqs) == len(chunks)
}


func main() {
  println(validate_single("aa aa bb cca"))
  println(validate_single("aa bb cca"))
  file, _ := os.Open("passphrases")
  reader := bufio.NewReader(file)
  passphrases := make([]string, 0, 0)
  for {
    line, err := reader.ReadString('\n')
    line = strings.TrimSpace(line)
    passphrases = append(passphrases, line)
    if err != nil {
      if err != io.EOF {
          fmt.Printf(" > Failed!: %v\n", err)
      }
      break
    }
  }
  println("LEN:", len(passphrases))
  println("RESULT", validate(passphrases))
}