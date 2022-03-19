from typing import List
import pytest

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        dp = [1 for _ in nums]
        
        for i in range(1,len(nums)):
            for j in range(i):
                if nums[i]>nums[j]:
                    dp[i]=max(dp[i], dp[j]+1)

        result = 0
        for v in dp:
            result=max(result,v)

        return result

@pytest.mark.parametrize("nums, want", [
    ([10,9,2,5,3,7,101,18], 4),
])
def test_lengOfLIS(nums, want):
    assert Solution().lengthOfLIS(nums) == want

