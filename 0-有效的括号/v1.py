class Solution:
    def isValid(self, s: str) -> bool:
        basic = {')': '(', '}': '{', ']': '['}
        stack = []
        for item in s:
            if not item in basic:
                stack.append(item)
            elif not stack or basic[item] != stack.pop():
                return False
        return True


if __name__ == '__main__':
    obj = Solution()
    ret1 = obj.isValid('[]{}()')
    ret2 = obj.isValid('{{[()]}}')
    ret3 = obj.isValid('(]')
    print(ret1, ret2, ret3)
