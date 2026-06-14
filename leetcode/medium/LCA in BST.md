# problem
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

# example
```python
**Input:** root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
**Output:** 6
**Explanation:** The LCA of nodes 2 and 8 is 6.
```

![[Pasted image 20260409155427.png]]
# 풀이
2 <= num of nodes <= 10^5
node 는 unique

bst
left < node < right

   2
1     3

time : O(H) (O(logN) ~ O(N)) / space : O(1)

```python
# Definition for a binary tree node.

class TreeNode:pi
	def __init__(self, x):
		self.val = x
		self.left = None
		self.right = None
		
		
		
def solve(root:TreeNode, p:TreeNode, q:TreeNode) -> TreeNode:
	while root:
		if root.val < p.val and root.val < q.val:
			root = root.right
		elif  root.val > p.val and root.val > q.val:
			root = root.left
		else:
			return root
	
```
p : 2, q : 8
root : 6
p: 2 / q: 4