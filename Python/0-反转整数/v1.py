class Solution:
    def reverse(self, x: int) -> int:
        s = str(x)
        if x < 0:
            ret = int(s[0] + s[-1:0:-1])
        else:
            ret = int(s[::-1])
        if not (-2**31 <= ret and ret <= 2**31-1):
            return 0
        return ret


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.reverse(123)
    ret2 = obj.reverse(120)
    ret3 = obj.reverse(-120)
    ret4 = obj.reverse(2**31)
    print(ret1, ret2, ret3, ret4)
