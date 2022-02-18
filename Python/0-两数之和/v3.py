# 一遍哈希表：核心想法--不再专门生成一张哈希表，而是一边往下遍历，一边查看之前的哈希表是否存在相应元素，否则将它存到哈希表中

# 时间复杂度：O(n)，每次查找只花费O(1)的时间，但总体上遍历了长度为n的列表
# 空间复杂度：O(n)，所需的额外空间取决于哈希表中存储的元素数量


class Solution:
    def twoSum(self, nums, target):
        basic = {}
        for index in range(len(nums)):
            ret = target - nums[index]
            if ret in basic.keys():
                return [basic.get(ret), index]
            basic[nums[index]] = index
