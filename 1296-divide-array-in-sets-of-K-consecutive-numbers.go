// Link: https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/

import (
    "container/heap"
)

type PQueue []int

func (pq PQueue) Len() int {
    return len(pq)
}

func (pq PQueue) Less(i,j int) bool {
    return pq[i] < pq[j]
} 

func (pq PQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
    *pq = append(*pq, x.(int))
}

func (pq *PQueue) Pop() interface{} {
    n := len(*pq)
    res := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return res
}

func isPossibleDivide(nums []int, k int) bool {
    n := len(nums)
    if n % k !=0 {
        return false
    }
    pq := make(PQueue, 0)
    hmap := make(map[int]int)
    heap.Init(&pq)
    for _, t := range nums {
        heap.Push(&pq, t)
        if _, ok := hmap[t]; !ok {
            hmap[t] = 1
        } else {
            hmap[t]++
        }
    }
    for pq.Len() > 0 {
        front := heap.Pop(&pq).(int)
        if hmap[front] <= 0 {
            continue
        }
        hmap[front]--
        for i:=1; i<k; i++ {
            if _, ok := hmap[front+i]; ok {
                if hmap[front+i] <= 0 {
                    return false
                }
                hmap[front+i]--
            } else {
                return false
            }
        }
    }
    return true
}

// Time complexity: O(NlogN)
// Space complexity: O(N)