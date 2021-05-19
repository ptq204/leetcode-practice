// Link: https://leetcode.com/problems/network-delay-time/

type Node struct {
    Key int
    Time int
}

type PriorityQueue []Node

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
    n := len(*pq)
    node := (*pq)[n-1]
    (*pq) = (*pq)[:n-1]
    return node
}

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].Time < pq[j].Time
}


func (pq PriorityQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
}


func networkDelayTime(times [][]int, n int, k int) int {
    graph := make([][]Node, n+1)
    for i:=0; i<n; i++ {
        graph[i] = make([]Node, 0)
    }
    for _, t := range times {
        from := t[0]
        to := t[1]
        time := t[2]
        graph[from] = append(graph[from], Node{to, time})
    }
    dist := make([]int, n+1)
    check := make([]bool, n+1)
    pq := make(PriorityQueue, 0)
    for i:=1; i<=n; i++ {
        dist[i] = math.MaxInt64
    }
    dist[k] = 0
    heap.Push(&pq, Node{k, 0})
    for pq.Len() > 0 {
        k := (heap.Pop(&pq).(Node)).Key
        check[k] = true
        for _, v := range graph[k] {
            travelThroughK := dist[k] + v.Time
            if !check[v.Key] && dist[v.Key] > travelThroughK {
                dist[v.Key] = dist[k] + v.Time
                heap.Push(&pq, Node{v.Key, travelThroughK})
            }
        }
    }
    ans := math.MinInt64
    for i:=1; i<=n; i++ {
        if dist[i] == math.MaxInt64 {
            return -1
        }
        if dist[i] > ans {
            ans = dist[i]
        }
    }
    return ans
}

// Time complexity: O(N*(logN + E))
// Space complexity: O(N)