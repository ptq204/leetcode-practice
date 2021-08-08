// Link: https://leetcode.com/problems/task-scheduler/

import (
    "container/heap"
)

// SOLUTION USING PRIORITY QUEUE AND TMP QUEUE
type PQueue []int

func (pq *PQueue) Push(x interface{}) {
    *pq = append(*pq, x.(int))
}

func (pq *PQueue) Pop() interface{} {
    n := len(*pq)
    task := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return task
}

func (pq PQueue) Less(i,j int) bool {
    return pq[i] > pq[j]
}

func (pq PQueue) Len() int {
    return len(pq)
}

func (pq PQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func leastInterval(tasks []byte, n int) int {
    hmap := make(map[byte]int)
    for _, t := range tasks {
        if _, ok := hmap[t]; !ok {
            hmap[t] = 1
        } else {
            hmap[t]++
        }
    }
    pq := make(PQueue, 0)
    heap.Init(&pq)
    buff := make([]int, 0)
    for _, v := range hmap {
        heap.Push(&pq, v)
    }
    ans := 0
    for pq.Len() > 0 {
        k := n + 1
        for k > 0 && pq.Len() > 0 {
            task := heap.Pop(&pq).(int)
            buff = append(buff, task-1)
            ans++
            k--
        }
        for len(buff) > 0 {
            task := buff[0]
            buff = buff[1:]
            if task > 0 {
                heap.Push(&pq, task)
            }
        }
        if pq.Len() == 0 {
            break
        }
        ans += k 
    }
    return ans
}

// Time complexity: O(NlogN)
// Space complexity: O(N)
