package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type numStack []int

func (stack numStack) Push(v int) numStack {
  return append(stack, v)
}

func (stack numStack) Pop() (numStack, int) {
  l := len(stack)
  return stack[:l-1], stack[l-1]
}

type opStack []byte

func (stack opStack) Push(v byte) opStack {
  return append(stack, v)
}

func (stack opStack) Pop() (opStack, byte) {
  l := len(stack)
  return stack[:l-1], stack[l-1]
}
func (stack opStack) Top() byte {
  if len(stack) == 0 {
    return 0
  }
  l := len(stack)
  return stack[l-1]
}

func ReadLines() []string {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  arr := make([]string, 0)
  for scanner.Scan() {
    arr = append(arr, scanner.Text())
  }
  return arr
}
func findIndex(str string) int {
  count := 0
  for k, v := range str {
    if v == '(' {
      count++
    } else if v == ')' {
      count--
    }
    if count == 0 {
      return k
    }
  }
  return -1
}
func evaluateExpression(op1 int, op2 int, operator byte) int {
  if operator == '*' {
    return op1 * op2
  } else if operator == '+' {
    return op1 + op2
  }
  return 0
}
func evaluateString(str string) int {
  var numstack numStack
  var opstack opStack
  var operator byte
  var op1, op2 int
  numStr := ""
  str += " "
  for k := 0; k < len(str); k++ {
    v := str[k]
    if v == ' ' {
      if numStr != "" {
        conv, _ := strconv.Atoi(numStr)
        numstack = numstack.Push(conv)
        numStr = ""
      }
    } else if v == '*' || v == '+' {
      top := opstack.Top()
      if v == '*' && top == '+' {
        opstack, _ = opstack.Pop()
        numstack, op1 = numstack.Pop()
        numstack, op2 = numstack.Pop()
        numstack = numstack.Push(evaluateExpression(op1, op2, top))
        k--
      } else {
        opstack = opstack.Push(v)
      }
    } else if v == '(' {
      index := findIndex(str[k:]) + k
      numstack = numstack.Push(evaluateString(str[k+1 : index]))
      k += index - k
    } else {
      numStr += string(v)
    }
  }
  for len(numstack) > 1 {
    numstack, op1 = numstack.Pop()
    numstack, op2 = numstack.Pop()
    opstack, operator = opstack.Pop()
    numstack = numstack.Push(evaluateExpression(op1, op2, operator))
  }
  return numstack[0]
}

func main() {
  data := ReadLines()
  sum := 0
  for _, v := range data {
    r := evaluateString(v)
    sum += r
  }
  fmt.Println(sum)

}
