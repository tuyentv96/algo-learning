from typing import List
import pytest
import 

class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        sub = [nums[0]]
        for num in nums[1:]:
            i = self.bs_leftmost(sub, num)
            if num > sub[-1]:
                sub.append(num)
            else:
                sub[i] = num
        return len(sub)

    def bs_leftmost(self, nums: List[int], num: int) -> int:
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] < num:
                lo = mid +1
            else:
                hi = mid

        return lo

    def bs_rightmost(self, nums: List[int], num: int) -> int:
        lo, hi = 0, len(nums)
        while lo < hi:
            mid = lo + (hi - lo) // 2
            if nums[mid] <= num:
                lo = mid +1
            else:
                hi = mid

        return lo

@pytest.mark.parametrize("nums, want", [
    ([0,1,0,3,2,3], 4),
    ([10,9,2,5,3,7,101,18], 4),
    ([3,5,6,2,5,4,19,5,6,7,12], 6),
])
def test_lengOfLIS(nums, want):
    assert Solution().lengthOfLIS(nums) == want

@pytest.mark.parametrize("nums, num, want", [
    ([1,2,3,4,5,6,7,8,9,10], 8, 7),
    ([1,2,3,4,5,6,7,8,9,10], 7, 6),
    ([2,5,6], 4, 1),
    ([2,5,6], 7, 2),
    ([0,2,2,2,4,4,5], 2, 1),
])
def test_bs_leftmost(nums, num, want):
    assert Solution().bs_leftmost(nums, num) == want

