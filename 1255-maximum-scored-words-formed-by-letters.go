// Link: https://leetcode.com/problems/maximum-score-words-formed-by-letters/

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func valid(word string, freq map[byte]int) bool {
    newFreq := make(map[byte]int)
    for i, _ := range word {
        if _, ok := newFreq[word[i]]; !ok {
            newFreq[word[i]] = 1
        } else {
            newFreq[word[i]]++
        }
    }
    for k, _ := range newFreq {
        if _, ok := freq[k]; !ok {
            return false
        }
        if freq[k] < newFreq[k] {
            return false
        }
    }
    return true
}

func calScore(word string, score []int) int {
    s := 0
    for _, c := range word {
        s += score[c-97]
    }
    return s
}

func solve(idx, n int, words []string, freq map[byte]int, score map[string]int) int {
    if idx >= n {
        return 0
    }
    res := 0
    for i:=idx; i<n; i++ {
        curr := words[i]
        if valid(curr, freq) {
            newFreq := make(map[byte]int)
            for k, v := range freq {
                newFreq[k] = v
            }
            for k, _ := range curr {
                newFreq[curr[k]]--
            }
            res = max(res, score[curr] + solve(i+1, n, words, newFreq, score))
        }
    }
    return res
}

func maxScoreWords(words []string, letters []byte, score []int) int {
    freq := make(map[byte]int)
    scoreMap := make(map[string]int)
    for _, l := range letters {
        if _, ok := freq[l]; !ok {
            freq[l] = 1
        } else {
            freq[l]++
        }
    }
    validSet := make([]string, 0)
    for _, w := range words {
        if valid(w, freq) {
            scoreMap[w] = calScore(w, score)
            validSet = append(validSet, w)
        }
    }
    return solve(0, len(validSet), validSet, freq, scoreMap)
}

// Time complexity: O(N^2) where N is number of words
// Space complexity: O(N)