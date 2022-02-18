class Solution:
    def removeDuplicates(self, nums: list) -> int:
        i = 0
        for j in range(len(nums)-1):
            if nums[j] != nums[j+1]:
                nums[i+1] = nums[j+1]
                i += 1
        del nums[i+1:]
        return len(nums)


if __name__ == '__main__':
    obj = Solution()
    ret = obj.removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])
    print(ret)
