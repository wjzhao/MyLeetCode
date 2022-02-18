class Solution:
    def isPalindrome(self, x: int) -> bool:
        s1 = str(x)
        s2 = s1[::-1]
        if s1 == s2:
            return True
        else:
            return False


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.isPalindrome(121)
    ret2 = obj.isPalindrome(-121)
    ret3 = obj.isPalindrome(10)
    print(ret1, ret2, ret3)
