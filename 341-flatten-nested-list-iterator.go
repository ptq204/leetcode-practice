// Link: https://leetcode.com/problems/flatten-nested-list-iterator/

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
    nums []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    nums := make([]*NestedInteger, 0)
    for i:=len(nestedList)-1; i>=0; i-- {
        nums = append(nums, nestedList[i])
    }
    return &NestedIterator{nums}
}

func (this *NestedIterator) Next() int {
    n := len(this.nums)
    res := this.nums[n-1].GetInteger()
    this.nums = this.nums[:n-1]
    return res
}

func (this *NestedIterator) HasNext() bool {
    n := len(this.nums)
    for n > 0 && !this.nums[n-1].IsInteger() {
        top := this.nums[n-1]
        this.nums = this.nums[:n-1]
        nestedList := top.GetList()
        for i:=len(nestedList)-1; i>=0; i-- {
            this.nums = append(this.nums, nestedList[i])
        }
        n = len(this.nums)
    }
    return len(this.nums) > 0
}

// Time complexity: O()
// Space complexity: O(N)