// Link: https://leetcode.com/problems/split-array-into-consecutive-subsequences/

import (
    "container/heap"
)

type PriorityQueue []int

func (pq *PriorityQueue) Len() int {
    return len(*pq)
}

func (pq *PriorityQueue) Swap(i,j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Less(i, j int) bool {
    return (*pq)[i] < (*pq)[j]
}

func (pq *PriorityQueue) Push(k interface{}) {
    *pq = append(*pq, k.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    k := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return k
}

func isPossible(nums []int) bool {
    hmap := make(map[int]*PriorityQueue)
    for _, k := range nums {
        if pq, ok := hmap[k-1]; ok && len(*pq) > 0 {
            minLen := heap.Pop(pq).(int)
            if _, ok := hmap[k]; !ok {
                hmap[k] = &PriorityQueue{}
            }
            q, _ := hmap[k]
            heap.Push(q, minLen + 1)
        } else {
            if _, ok := hmap[k]; !ok {
                hmap[k] = &PriorityQueue{}
            }
            q, _ := hmap[k]
            heap.Push(q, 1)
        }
    }
    for _, v := range hmap {
        if len(*v) == 0 {
            continue
        }
        if heap.Pop(v).(int) < 3 {
            return false
        }
    }
    return true
}

// Time complexity: O(NlogN)
// Space complexity: O(N)