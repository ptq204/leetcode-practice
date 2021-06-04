// Link:

import (
    "container/heap"
)

// SOLUTION USING TRIE + HEAP
type PQueue []string

func (pq PQueue) Len() int {
    return len(pq)
}

func (pq PQueue) Swap(i,j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
    *pq = append(*pq, x.(string))
}

func (pq *PQueue) Pop() interface{} {
    n := len(*pq)
    tmp := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return tmp
}

func (pq PQueue) Less(i,j int) bool {
    return pq[i] > pq[j]
}

type Trie struct {
    Next map[rune]*Trie
    Res PQueue
}

func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Slice(products, func (i,j int) bool {
        return products[i] < products[j]
    })
    root := &Trie{make(map[rune]*Trie), make(PQueue, 0)}
    for _, p := range products {
        curr := root
        for _, c := range p {
            if _, ok := curr.Next[c]; !ok {
                pq := make(PQueue, 0)
                heap.Init(&pq)
                curr.Next[c] = &Trie{make(map[rune]*Trie), pq}
            }
            curr = curr.Next[c]
            suggestions := &curr.Res
            heap.Push(suggestions, p)
            if suggestions.Len() > 3 {
                heap.Pop(suggestions)
            }
            sort.Slice(curr.Res, func (i,j int) bool {
                return curr.Res[i] < curr.Res[j]
            })
        }
    }
    ans := make([][]string, 0)
    curr := root
    for i, c := range searchWord {
        if curr.Next[c] == nil {
            for j:=i; j<len(searchWord); j++ {
                ans = append(ans, []string{})
            }
            return ans
        }
        curr = curr.Next[c]
        ans = append(ans, curr.Res)
    }
    return ans
}
// N is number of words, S is medium length of words, K is 3
// Time complexity: O(N*S*logK)
// Space complexity: O(N*S*K)


// SOLUTION USING TRIE + SORT
type Trie struct {
    Next map[rune]*Trie
    Res []string
}

func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Slice(products, func (i,j int) bool {
        return products[i] < products[j]
    })
    root := &Trie{make(map[rune]*Trie), make([]string, 0)}
    for _, p := range products {
        curr := root
        for _, c := range p {
            if _, ok := curr.Next[c]; !ok {
                curr.Next[c] = &Trie{make(map[rune]*Trie), make([]string, 0)}
            }
            suggestions := curr.Next[c].Res
            if len(suggestions) < 3 {
                suggestions = append(suggestions, p)
            }
            curr.Next[c].Res = suggestions
            curr = curr.Next[c]
        }
    }
    ans := make([][]string, 0)
    curr := root
    for i, c := range searchWord {
        if curr.Next[c] == nil {
            for j:=i; j<len(searchWord); j++ {
                ans = append(ans, []string{})
            }
            return ans
        }
        curr = curr.Next[c]
        ans = append(ans, curr.Res)
    }
    return ans
}

// Time complexity: O(NlogN)
// Space complexity: O(O*N*K)