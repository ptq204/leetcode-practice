// Link: https://leetcode.com/problems/open-the-lock/

func openLock(deadends []string, target string) int {
    start := "0000"
    if target == start {
        return 0
    }
    hmap := make(map[string]bool)
    for _, d := range deadends {
        hmap[d] = true
    }
    if _, ok := hmap[start]; ok {
        return -1
    }
    queue := make([]string, 0)
    queue = append(queue, start)
    level := 0
    
    for len(queue) > 0 {
        size := len(queue)
        for size > 0 {
            front := queue[0]
            if front == target {
                return level
            }
            queue = queue[1:]
            for i:=0; i<4; i++ {
                prev := int(front[i])-49
                next := int(front[i])-47
                if prev < 0 {
                    prev = 9
                }
                if next > 9 {
                    next = 0
                }
                key1 := front[:i] + strconv.Itoa(prev) + front[i+1:]
                if _, ok := hmap[key1]; !ok {
                    queue = append(queue, key1)
                    hmap[key1] = true
                }
                key2 := front[:i] + strconv.Itoa(next) + front[i+1:]
                if _, ok := hmap[key2]; !ok {
                    queue = append(queue, key2)
                    hmap[key2] = true
                }
            }
            size--
        }
        level++
    }
    return -1
}

// Time complexity: 
// Space complexity: 