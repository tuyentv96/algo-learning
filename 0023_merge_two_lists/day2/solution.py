from queue import PriorityQueue
from typing import Optional, List

# Definition for singly-linked list.
class ListNode:
    ListNode.__lt__ = lambda self, other: self.val < other.val
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        head = point = ListNode(0)
        q = PriorityQueue()
        for l in lists:
            if l:
                q.put((l.val,l))

        while not q.empty():
            val, node = q.get()
            point.next = node
            point = point.next
            node = node.next
            if node:
                q.put((node.val,node))

        return head.next

