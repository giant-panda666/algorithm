#!/usr/bin/python

#Given an array and a value, remove all instances of that value in place and return the new length.
#
#Do not allocate extra space for another array, you must do this in place with constant memory.
#
#The order of elements can be changed. It doesn't matter what you leave beyond the new length.
#
#Example:
#Given input array nums = [3,2,2,3], val = 3
#
#Your function should return length = 2, with the first two elements of nums being 2.

class Solution(object):
    def removeElement(self, nums, val):
        """
        :type nums: List[int]
        :type val: int
        :rtype int
        """
        i, j = 0, len(nums)-1
        while j >= i:
            while i < len(nums) and nums[i] != val:
                i += 1
            while j >= 0 and nums[j] == val:
                j -= 1
            if j >= i:
                nums[i], nums[j] = nums[j], nums[i]
        nums = nums[0:j+1]
        return j+1

s = Solution()
a1 = [1]
print s.removeElement(a1, 1)
