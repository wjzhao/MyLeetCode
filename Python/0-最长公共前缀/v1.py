class Solution:
    def longestCommonPrefix(self, strs: list) -> str:
        if len(strs) == 0:
            return ""
        index = 0
        min_length = len(strs[0])
        for i in range(len(strs)):
            if len(strs[i]) < min_length:
                min_length = len(strs[i])
                index = i

        for j in range(len(strs[index])):
            for k in range(len(strs)):
                if strs[index][j] != strs[k][j]:
                    return strs[index][:j]
        return strs[index]


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.longestCommonPrefix(["flower", "flow", "flight"])
    ret2 = obj.longestCommonPrefix(["dog", "racecar", "car"])
    print(ret1, ret2)
