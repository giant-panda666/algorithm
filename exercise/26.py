#!/usr/bin/python

#Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
#
#Do not allocate extra space for another array, you must do this in place with constant memory.
#
#For example,
#Given input array nums = [1,1,2],
#
#Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 0:
            return 0
        while True:
            i, j = 0, 1
            while j < len(nums):
                if nums[i] == nums[j]:
                    del nums[i]
                    break
                else:
                    i += 1
                    j += 1
            if j == len(nums):
                return len(nums)

s = Solution()
l = [1, 1, 2]
print s.removeDuplicates(l)

l = [1, 2, 2, 3, 4, 4]
print s.removeDuplicates(l)
