#!/usr/bin/python

#Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
#
#For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
#
#Note:
#    1.You must do this in-place without making a copy of the array.
#    2.Minimize the total number of operations.

class Solution(object):
    def moveZeroes(self, nums):
        i, j = 0, 0
        while j < len(nums):
            if nums[j] != 0:
                if j != i:
                    nums[i], nums[j] = nums[j], nums[i]
                i += 1
            j += 1
        print nums

s = Solution()
nums = [0, 1, 0, 3, 12]
print s.moveZeroes(nums)

nums = []
print s.moveZeroes(nums)

nums = [1, 2, 0, 4]
print s.moveZeroes(nums)
