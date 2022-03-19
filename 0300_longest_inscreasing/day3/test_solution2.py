from typing import List
import pytest

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        sub = [nums[0]]

        for num in nums[1:]:
            if num > sub[-1]:
                sub.append(num)
            else:
                i = 0
                while sub[i] < num:
                    i += 1

                sub[i] = num

        return len(sub)

@pytest.mark.parametrize("nums, want", [
    ([0,1,0,3,2,3], 4),
    ([10,9,2,5,3,7,101,18], 4),
])
def test_lengOfLIS(nums, want):
    assert Solution().lengthOfLIS(nums) == want

