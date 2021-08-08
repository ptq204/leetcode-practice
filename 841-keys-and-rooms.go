// Link: https://leetcode.com/problems/keys-and-rooms/

func canVisitAllRooms(rooms [][]int) bool {
    check := make([]bool, len(rooms))
    queue := make([]int, 0)
    for _, k := range rooms[0] {
        queue = append(queue, k)
    }
    check[0] = true
    for len(queue) > 0 {
        front := queue[0]
        queue = queue[1:]
        if check[front] {
            continue
        }
        check[front] = true
        for _, k := range rooms[front] {
            queue = append(queue, k)
        }
    }
    for _, c := range check {
        if !c {
            return false
        }
    }
    return true
}

// Time complexity: O(N)
// Space complexity: O(N)
