# 按照ascii码大小来排序


class Solution:
    def longestCommonPrefix(self, strs: list) -> str:
        if not strs: return ""
        s1 = min(strs)
        s2 = max(strs)
        for i, v in enumerate(s1):
            if v != s2[i]:
                return s2[:i]
        return s1


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.longestCommonPrefix(["flower", "flow", "flight"])
    ret2 = obj.longestCommonPrefix(["dog", "racecar", "car"])
    print(ret1, ret2)
