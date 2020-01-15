package main

import "fmt"
import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
 }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  var result *ListNode  
  var resultStepper *ListNode
  var carryOver int

  for l1 != nil || l2 != nil {
    var sum int  

    if l1 != nil {
      sum += l1.Val
      l1= l1.Next
    }

    if l2 != nil {
      sum += l2.Val
      l2 = l2.Next
    }

    sum+=carryOver
    val := int(math.Mod(float64(sum), 10))
    carryOver = sum/10

    ls := ListNode{val, nil}

    if result == nil {
      result = &ls
      resultStepper = &ls
    } else {
      resultStepper.Next = &ls
      resultStepper = resultStepper.Next
    }
    
  }

  if carryOver != 0 {
    ls := ListNode{carryOver, nil}
    resultStepper.Next = &ls
  }

  return result
}

func main() {
  fmt.Println("Hello World")
}