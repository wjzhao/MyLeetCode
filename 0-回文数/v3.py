class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        for i in range(len(s)):
            for j in range(len(s)-i-1, -1, -1):
                if s[i] == s[j]:
                    break
                else:
                    return False
        return True


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.isPalindrome(120)
    ret2 = obj.isPalindrome(-120)
    ret3 = obj.isPalindrome(121)
    print(ret1, ret2, ret3)
