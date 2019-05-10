class Solution:
    def isPalindrome(self, x: int) -> bool:
        temp = []
        ret = []
        s = str(x)
        for elem in s:
            temp.append(elem)
        for i in range(len(temp)):
            ret.append(temp.pop())
        result = ''.join(ret)
        if result == s:
            return True
        else:
            return False


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.isPalindrome(120)
    ret2 = obj.isPalindrome(121)
    ret3 = obj.isPalindrome(-12)
    print(ret1, ret2, ret3)
