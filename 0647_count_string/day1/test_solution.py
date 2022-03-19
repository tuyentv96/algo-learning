from solution import Solution
import pytest

@pytest.mark.parametrize("s, want", [
    ("abc", 3),
    ("aaa", 6),
])
def test(s, want):
    assert Solution().countSubstrings(s) == want

