// Link: https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    a := 0
    b := 1
    ans := a + b
    for i:=2; i<=n; i++ {
        ans = a + b
        a = b
        b = ans
    }
    return ans
}