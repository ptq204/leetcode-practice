// Link:

import (
    "container/heap"
)

type Item struct {
    Index int
    Val int
}

type PQueue []Item

func (pq *PQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Item))
}

func (pq PQueue) Peek() interface{} {
    return pq[0]
}

func (pq *PQueue) Pop() interface{} {
    n := len(*pq)
    item := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return item
}

func (pq PQueue) Less(i,j int) bool {
    return pq[i].Val < pq[j].Val
}

func (pq PQueue) Len() int {
    return len(pq)
}

func (pq PQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func dailyTemperatures(T []int) []int {
    if len(T) == 0 {
        return []int{}
    }
    pq := make(PQueue, 0)
    heap.Init(&pq)
    ans := make([]int, len(T))
    
    for i, t := range T {
        for pq.Len() > 0 {
            front := pq.Peek().(Item)
            if front.Val < t {
                ans[front.Index] = i-front.Index
                heap.Pop(&pq)
            } else {
                break
            }
        }
        heap.Push(&pq, Item{i, t})
    }
    return ans
}

// Time complexity: O(NlogN)
// Space complexity: O(N)