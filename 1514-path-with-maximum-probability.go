// Link:

import (
    "container/heap"
    "fmt"
)

type Node struct {
    Key int
    Prob float64
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    node := (*pq)[n-1]
    (*pq) = (*pq)[:n-1]
    return node
}

func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].Prob > pq[j].Prob
}

func (pq PriorityQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
} 

func getMaxNextDistIdx(dist []float64, check []bool) int {
    maxDist := -1.0
    idx := -1
    for i:=0; i<len(dist); i++ {
        if dist[i] > maxDist && !check[i] {
            maxDist = dist[i]
            idx = i
        }
    }
    return idx
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    graph := make([][]Node, n)
    for i:=0; i<n; i++ {
        graph[i] = make([]Node, 0)
    }
    for i:=0; i<len(edges); i++ {
        from := edges[i][0]
        to := edges[i][1]
        graph[from] = append(graph[from], Node{to, succProb[i]})
        graph[to] = append(graph[to], Node{from, succProb[i]})
    }

    dist := make([]float64, n)
    check := make([]bool, n)
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    for i:=0; i<n; i++ {
        dist[i] = -1.0
    }
    dist[start] = 1.0
    heap.Push(&pq, Node{start, 1})

    for pq.Len() > 0 {
        k := (heap.Pop(&pq).(Node)).Key
        check[k] = true
        if k == end {
            break
        }
        for _, v := range graph[k] {
            if dist[k] * v.Prob > dist[v.Key] && !check[v.Key] {
                dist[v.Key] = dist[k] * v.Prob
                heap.Push(&pq, Node{v.Key, dist[k] * v.Prob})
            }
        }
    }
    if dist[end] == -1.0 {
        return 0.0
    }
    return dist[end]
}

// Time complexity: O(V*(logV+E))
// Space complexity: O(N)