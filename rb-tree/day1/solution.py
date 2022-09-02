class Node:
    def __init__(self, val) -> None:
        self.val = val
        self.parent = None
        self.left = None
        self.right = None
        self.red = True

class RBTree:
    def __init__(self) -> None:
        self.root = None

    def insert(self, node: Node):
        if not self.root:
            self.root = node 
            return

        prev = None
        tmp = self.root

        while tmp:
            prev = tmp
            if val < tmp.val:
                tmp = tmp.left
            else:
                tmp = tmp.right

        node.parent = prev
        if node.val < prev.val:
            prev.left = node
        else:
            prev.right = node

    def rebalance(self):
        
