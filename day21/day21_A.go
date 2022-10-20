package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
)

type Item struct {
  foods     []string
  allergens []string
}

func ReadLines() []Item {
  f, _ := os.Open("input.txt")
  defer f.Close()
  scanner := bufio.NewScanner(f)
  arr := make([]Item, 0)
  for scanner.Scan() {
    arr = append(arr, Item{})
    text := scanner.Text()
    reg := regexp.MustCompile(`(?m)\bcontains\s+(.*)$`)
    allergens := reg.FindString(text)
    allergens = allergens[9:len(allergens)-1] + " "
    str := ""
    for i := 0; i < len(allergens); i++ {
      if allergens[i] == ',' || allergens[i] == ' ' {
        arr[len(arr)-1].allergens = append(arr[len(arr)-1].allergens, str)
        str = ""
        i++
        continue
      }
      str += string(allergens[i])
    }
    str = ""
    for _, v := range text {
      if v == '(' {
        break
      }
      if v == ' ' {
        arr[len(arr)-1].foods = append(arr[len(arr)-1].foods, str)
        str = ""
        continue
      }
      str += string(v)
    }
  }
  return arr
}

/*
Strategy:

 1. Compare two Item's :
    1a. Get intersection of their "foods".
    1b. Gen intersection of their "allergens".
    1c. If both sets are non-empty, create a new item from those intersection of foods and allergens. Delete the intersection items from the original Items lists.
 2. Do Step 1 for all possible pairs of items until there is no more new Item emerging.
 3. Look at the Items with non-empty allergen list and the remove the elements of their foods list from other Items foods list with empty allergens list.
 4. Count the items with empty allergen-list.
*/

func intersection(list1 []string, list2 []string) (arr []string) {
  for _, l1 := range list1 {
    for _, l2 := range list2 {
      if l1 == l2 {
        arr = append(arr, l1)
      }
    }
  }
  return arr
}
func removeElems(orig []string, elems []string) (arr []string) {
  flag := true
  for _, v := range orig {
    flag = true
    for _, elem := range elems {
      if v == elem {
        flag = false
        break
      }
    }
    if flag {
      arr = append(arr, v)
    }
  }
  return
}
func compare(i1 Item, i2 Item) (bool, Item) {
  foodsIntersection := intersection(i1.foods, i2.foods)
  allergensIntersection := intersection(i1.allergens, i2.allergens)
  if len(foodsIntersection) == 0 || len(allergensIntersection) == 0 {
    return false, Item{}
  } else {
    return true, Item{foods: foodsIntersection, allergens: allergensIntersection}
  }
}
func removeIntersections(item *Item, intersection Item) {
  (*item).allergens = removeElems((item.allergens), intersection.allergens)
  (*item).foods = removeElems((item.foods), intersection.foods)
}
func execute(items []Item) []Item {

  for i := 0; i < len(items)-1; i++ {
    for j := i + 1; j < len(items); j++ {
      ok, intersection := compare(items[i], items[j])
      if ok {
        removeIntersections(&items[i], intersection)
        removeIntersections(&items[j], intersection)
        items = append(items, intersection)
      }
    }
  }
  return items
}
func deleteStr(items []Item, str string) []Item {
  for k := range items {
    if len(items[k].allergens) != 0 {
      continue
    }
    for j, v := range items[k].foods {
      if v == str {
        last := len(items[k].foods) - 1
        items[k].foods[j] = items[k].foods[last]
        items[k].foods = items[k].foods[:last]
      }
    }
  }
  return items
}
func deletePossible(items []Item) []Item {
  for _, v := range items {
    if len(v.allergens) > 0 {
      for _, str := range v.foods {
        items = deleteStr(items, str)
      }
    }
  }
  return items
}
func count(items []Item) (c int) {
  for _, v := range items {
    if len(v.allergens) == 0 {
      c += len(v.foods)
    }
  }
  return
}
func main() {
  data := ReadLines()
  for {
    length := len(data)
    data = execute(data)
    if length == len(data) {
      break
    }
  }
  data = deletePossible(data)
  count := count(data)
  for _, v := range data {
    fmt.Println(v)
  }
  fmt.Println(count)

}
