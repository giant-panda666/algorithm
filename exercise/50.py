#!/usr/bin/python

class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        if n == 0:
            return 1
        if n < 0:
            n = -n
            x = 1/x
        if n % 2 == 0 :
            return self.myPow(x*x, n/2)
        else:
            return x*self.myPow(x*x, n/2)

s = Solution()
print s.myPow(-2.0, 3)
print s.myPow(2.0, 4)
print s.myPow(2.0, -2)
print s.myPow(2.0, -3)
print s.myPow(0.0, 20)
print s.myPow(2.0, 0)
print s.myPow(2.00000, -2147483648)
