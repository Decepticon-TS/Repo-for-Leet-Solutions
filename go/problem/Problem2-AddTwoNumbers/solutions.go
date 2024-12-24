package main

import "fmt"

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Update carry and create a new node for the current digit
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	l1 := createLinkedList([]int{0})
	l2 := createLinkedList([]int{7, 3})
	result := addTwoNumbers(l1, l2)
	printLinkedList(result)

	// Example 2
	l1 = createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = createLinkedList([]int{9, 9, 9, 9})
	result = addTwoNumbers(l1, l2)
	printLinkedList(result)
}
