// Link: https://leetcode.com/problems/lru-cache/

type Node struct {
    Key int
    Val int
    Next *Node
    Prev *Node
}

type LRUCache struct {
    cache map[int]*Node
    capacity int
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    cache := make(map[int]*Node)
    return LRUCache{cache, capacity, nil, nil}
}

func (this *LRUCache) handleLRULogic(key int) {
    curr := this.cache[key]
    prev := curr.Prev
    next := curr.Next
    if curr == this.head {
        this.head = next
    }
    if prev != nil {
        prev.Next = curr.Next   
    }
    next.Prev = prev
    curr.Next = nil
    this.tail.Next = curr
    curr.Prev = this.tail
    this.tail = curr
    this.cache[key] = curr
}

func (this *LRUCache) Get(key int) int {
    if _, ok := this.cache[key]; !ok {
        return -1
    }
    res := this.cache[key].Val
    // if not LRU node -> handle LRU logic for middle node
    if this.cache[key] != this.tail {
        this.handleLRULogic(key)
    }
    return res
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; ok {
        if this.cache[key] == this.tail {
            this.cache[key].Val = value
            return
        }
        // handle LRU here
        this.handleLRULogic(key)
        this.cache[key].Val = value
        return
    }
    // If key not exist -> insert at last
    size := len(this.cache)
    if size == 0 {
        this.head = &Node{key, value, nil, nil}
        this.tail = this.head
    } else if size < this.capacity  {
        this.tail.Next = &Node{key, value, nil, this.tail}
        this.tail = this.tail.Next
    } else {
        delete(this.cache, this.head.Key)
        this.head = this.head.Next
        if this.head == nil {
            this.head = &Node{key, value, nil, nil}
            this.tail = this.head
        } else {
            this.tail.Next = &Node{key, value, nil, this.tail}
            this.tail = this.tail.Next
        }
    }
    this.cache[key] = this.tail
}


// Time complexity for GET and PUT: O(1)
// Space complexity: O(N)