# 暴力法，遍历每个元素 x，并找到是否存在一个值与 target-x 相等的目标元素

# 时间复杂度：O(n^2)，对于每个元素，试图通过遍历数组的其余部分来寻找它所对应的目标元素，这将耗费O(n)的时间，因此时间复杂度是O(n^2)
# 空间复杂度：O(1)


class Solution:
    def twoSum(self, nums, target):
        for i in range(len(nums)-1):
            for j in range(i+1, len(nums)):
                if nums[j] == target - nums[i]:
                    return [i, j]
