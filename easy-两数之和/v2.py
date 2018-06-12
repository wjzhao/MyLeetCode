# 两遍哈希表：核心想法--用哈希表保持数组中每个元素与其索引的相互对应关系

# 时间复杂度：O(n)，把含有n个元素的列表遍历了两次，由于哈希表将查找时间缩短到了O(1)，所以时间复杂度是O(n)
# 空间复杂度：O(n)，所需的额外空间取决于哈希表中存储的元素数量


class Solution:
    def twoSum(self, nums, target):
        basic = {}
        for index in range(len(nums)):
            basic[nums[index]] = index
        for i in range(len(nums)):
            ret = target - nums[i]
            if ret in basic.keys() and (basic.get(ret) != i):
                return [i, basic.get(ret)]

