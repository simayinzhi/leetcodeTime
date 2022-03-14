package code

import "fmt"

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BalanceTreeNode struct {
	Val   int
	Bf    int
	Left  *BalanceTreeNode
	Right *BalanceTreeNode
}

const defaultStackIncrementSize int = 10

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val)
		fmt.Print(",")
		list = list.Next
	}
	println("")
}

func CreateList(slice ...int) *ListNode {
	result := &ListNode{}
	node := result
	for i := 0; i < len(slice); i++ {
		node.Next = &ListNode{slice[i], nil}
		node = node.Next
	}
	return result.Next
}

type myStack struct {
	list   []interface{}
	length int
}

func InitMyStack(length int) myStack {
	return myStack{make([]interface{}, length), 0}
}
func InitDefaultLengthMyStack() myStack {
	return InitMyStack(defaultStackIncrementSize)
}

func (s myStack) Peek() interface{} {
	if s.length > 0 {
		return s.list[s.length-1]
	}
	return nil
}

func (s *myStack) Pop() interface{} {
	if s.length > 0 {
		top := s.list[s.length-1]
		s.list[s.length-1] = nil
		s.length--
		return top
	} else {
		return nil
	}
}
func (s *myStack) Push(str interface{}) {
	if s.length == len(s.list) {
		s.AddCapacity()
	}
	s.list[s.length] = str
	s.length++
}

func (s *myStack) AddCapacity() {
	list := make([]interface{}, len(s.list)+defaultStackIncrementSize)
	for i := 0; i < len(s.list); i++ {
		list[i] = s.list[i]
	}
	s.list = list
}

type PriorityQueue struct {
	capacity int
	dump     []int
}

func InitPriorityQueue(capacity int) PriorityQueue {
	return PriorityQueue{capacity, make([]int, 0, capacity)}
}

func (queue *PriorityQueue) InsertPriorityQueue(num int) bool {
	if queue.capacity == len(queue.dump) {
		return false
	}
	queue.dump = append(queue.dump, num)
	percolateUp(queue.dump)
	return true
}

func percolateUp(slice []int) {
	index := len(slice) - 1
	for index > 0 {
		parentIndex := findParentIndexByCurrentIndex(index)
		if slice[index] > slice[parentIndex] {
			SwapElements(slice, index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

func (queue *PriorityQueue) Size() int {
	return len(queue.dump)
}

func (queue *PriorityQueue) GetMax() int {
	return queue.dump[0]
}

func (queue *PriorityQueue) DelMax() int {
	if len(queue.dump) == 0 {
		return -9999
	}
	max := queue.dump[0]
	queue.dump[0] = queue.dump[len(queue.dump)-1]
	queue.dump = queue.dump[:len(queue.dump)-1]

	percolateDown(queue.dump, 0, len(queue.dump))
	return max
}

func percolateDown(slice []int, position int, length int) {
	for position*2+1 < length {
		maxChild := position*2 + 1
		if maxChild+1 < length {
			if slice[maxChild] < slice[maxChild+1] {
				maxChild = maxChild + 1
			}
		}
		if slice[position] < slice[maxChild] {
			SwapElements(slice, position, maxChild)
			position = maxChild
		} else {
			break
		}
	}
}
func (queue *PriorityQueue) empty() {
	queue.dump = queue.dump[:0]
}

func findParentIndexByCurrentIndex(index int) int {
	return (index - 1) / 2
}
func findLeftChildren(index int) int {
	return index*2 + 1
}
func findRightChildren(index int) int {
	return index*2 + 2
}

func SwapElements(slice []int, xIndex int, yIndex int) {
	if xIndex == yIndex {
		return
	}
	temp := slice[xIndex]
	slice[xIndex] = slice[yIndex]
	slice[yIndex] = temp
}

func Heapify(slice []int, low int, height int) []int {
	if len(slice) <= 1 {
		return slice
	}
	for i := height; i >= low; i-- {
		percolateDown(slice, i, height+1)
	}
	return slice
}

func HeapSort(slice []int, top int) []int {
	for i := 0; i < top; i++ {
		Heapify(slice[i:], 0, len(slice[i:])-1)
	}
	return slice[0:top]
}
