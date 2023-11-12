package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	current := list1
	for current.Next != nil {
		current = current.Next
	}
	current.Next = list2

	return mergeSort(list1)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return dummy.Next
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddle(head)

	left := head
	right := mid.Next
	mid.Next = nil
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
