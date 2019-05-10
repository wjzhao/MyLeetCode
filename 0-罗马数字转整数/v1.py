# 对字符串从左遍历，如果当前字符代表值不小于右边，就加上该值；否则减去该值


class Solution:
    def romanToInt(self, s: str) -> int:
        chars = {'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
        ret = 0
        for i in range(len(s)):
            if i < len(s)-1 and chars[s[i]] < chars[s[i+1]]:
                ret -= chars[s[i]]
            else:
                ret += chars[s[i]]
        return ret


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.romanToInt('III')
    ret2 = obj.romanToInt('LVIII')
    ret3 = obj.romanToInt('IX')
    print(ret1, ret2, ret3)
