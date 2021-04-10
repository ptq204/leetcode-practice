// Link: https://leetcode.com/problems/distant-barcodes/

// SOLUTION 1: using HashMap and if there always has an answer, we fill barcodes every 2 distance
func rearrangeBarcodes(barcodes []int) []int {
    hmap := make(map[int]int)
    maxBCFreq := 0
    mostBC := barcodes[0]
    ans := make([]int, len(barcodes))
    for _, bc := range barcodes {
        if _, ok := hmap[bc]; !ok {
            hmap[bc] = 1
        } else {
            hmap[bc]++
        }
        if hmap[bc] > maxBCFreq {
            maxBCFreq = hmap[bc]
            mostBC = bc
        }
    }
    i := 0
    idx := 0
    // The most frequent barcodes will appear no more than n/2 + 1
    for i < maxBCFreq && idx < len(barcodes) {
        ans[idx] = mostBC
        idx+=2
        i++
    }
    delete(hmap, mostBC)
    for k, v := range hmap {
        for j:=0; j<v; j++ {
            if idx >= len(barcodes) {
                idx = 1
            }
            ans[idx] = k
            idx += 2
        }
    }
    return ans
}
// Time complexity: O(N)
// Space complexity: O(N)

// SOLUTION 2: using Priority Queue
import (
    "container/heap"
)

type PriorityQueue [][]int

func (pq *PriorityQueue) Len() int {
    return len(*pq)
}

func (pq *PriorityQueue) Swap(i,j int) {
    (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Less(i, j int) bool {
    return (*pq)[i][1] > (*pq)[j][1]
}

func (pq *PriorityQueue) Push(k interface{}) {
    *pq = append(*pq, k.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    k := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return k
}

func rearrangeBarcodes(barcodes []int) []int {
    hmap := make(map[int]int)
    ans := make([]int, len(barcodes))
    for _, bc := range barcodes {
        if _, ok := hmap[bc]; !ok {
            hmap[bc] = 1
        } else {
            hmap[bc]++
        }
    }
    pq := &PriorityQueue{}
    heap.Init(pq)
    for k, v := range hmap {
        heap.Push(pq, []int{k, v})
    }
    idx := 0
    for len(*pq) > 0 {
        val := heap.Pop(pq)
        bcode := val.([]int) 
        for c:=0; c<bcode[1]; c++ {
            if idx >= len(barcodes) {
                idx = 1
            }
            ans[idx] = bcode[0]
            idx+=2
        }
    }
    return ans
}
