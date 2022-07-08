class Solution:
    def isBalanced(self, root):
        def depth(root):
            if not root:
                return 0
            return max(depth(root.left), depth(root.right)) + 1

        def solution(root):
            if not root:
                return True
            if abs(depth(root.left) - depth(root.right)) > 1:
                return False
            return solution(root.left) and solution(root.right)
        return solution(root)
        