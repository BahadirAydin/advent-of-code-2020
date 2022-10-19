package main

//IMPORTANT: You should add "mask = end" to the end of your input as a newline to run this code
import (
  "bufio"
  "fmt"
  "math"
  "os"
  "strconv"
  "strings"
)

type Operation struct {
  mask    string
  indexes []int
  values  []int64
}

func ReadLines() []Operation {

  f, _ := os.Open("input.txt")
  defer f.Close()
  arr := make([]Operation, 0)
  scanner := bufio.NewScanner(f)
  var tmp Operation
  indexes := make([]int, 0)
  values := make([]int64, 0)
  for scanner.Scan() {
    text := scanner.Text()
    if text[1] == 'a' {
      cpyIndex := make([]int, len(indexes))
      copy(cpyIndex, indexes)
      cpyValues := make([]int64, len(values))
      copy(cpyValues, values)
      tmp.indexes = cpyIndex
      tmp.values = cpyValues
      arr = append(arr, tmp)
      tmp.mask = text[7:]
      indexes = nil
      values = nil
      continue
    }
    rightInd := strings.Index(text, "]")
    index, _ := strconv.Atoi(text[4:rightInd])
    b, _ := strconv.Atoi(text[rightInd+4:])
    value := int64(b)
    indexes = append(indexes, index)
    values = append(values, value)
  }
  return arr[1:]
}
func applyMask(mask string, v int64) (num int) {

  s := strconv.FormatInt(v, 2)
  var count float64
  for i := 0; i < len(mask); i++ {
    if mask[35-i] == '1' {
      num += int(math.Pow(2, count))
    } else if mask[35-i] == '0' {
      count++
      continue
    } else if i < len(s) && s[len(s)-1-i] == '1' {
      num += int(math.Pow(2, count))
    }
    count++
  }
  return
}
func main() {
  data := ReadLines()
  space := make(map[int]int)
  for _, v := range data {
    mask := v.mask
    for i := 0; i < len(v.indexes); i++ {
      space[v.indexes[i]] = applyMask(mask, v.values[i])
    }
  }
  sum := 0
  for _, v := range space {
    sum += v
  }
  fmt.Println(sum)
}
