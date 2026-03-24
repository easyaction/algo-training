[link](https://leetcode.com/problems/merge-k-sorted-lists/description/)
### 문제 이해
k linked list가 주어질때, 이를 하나의 정렬된 linked list로 머지하라
각 리스트는 오름차순으로 정렬되어있음
### 추상화
모든 원소를 오름차순으로 머지해야함
 min-heap을 이용하여 머지 : O(NlogN)
### 계획
각 리스트를 순회하면서
	node의 val만 heapq에 넣고 다음 원소로 이동
heapq로 pop하면서 노드 엮어주기
### 코드
```python
import heapq

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # brute-force using min heap
        # time : O(Mlog(M)) space : O(M)
        # M = number of all nodes
        pq = []
        for head in lists:
            cur = head
            while cur:
                tmp = cur.next
                heapq.heappush(pq, (cur.val, )
                cur = tmp
        prev = None
        head = None
        while pq:
            node = ListNode(heapq.heappop(pq))
            if prev is None:
                prev = node
                head = prev
            else:
                prev.next = node
                prev = node
        return head
```
heap에 모든 원소 안넣게하여 공간 복잡도 절약
```python
import heapq

class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        # using min heap with minimum heap push 
        # time : O(Mlogk) space : O(k) 
        # M = number of all nodes
        pq = []
        for i, head in enumerate(lists):
            if head:
                heapq.heappush(pq, (head.val, i, head))
        dummy = ListNode(-1)
        cur = dummy
        while pq:
            val, i, node = heapq.heappop(pq)
            cur.next = node
            cur = node
            if node.next:
                heapq.heappush(pq, (node.next.val, i, node.next))
        return dummy.next
```
### 회고
모든 원소를 다 꺼냈다가 넣으면 공간복잡도가 불필요하게 커짐
	각 리스트의 커서에서 다음을 탐색할때마다 heappush하면 공간복잡도 kl -> k로 절악 가능
분할 정복으로도 풀수 있음