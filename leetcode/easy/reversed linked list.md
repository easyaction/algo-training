[link](https://leetcode.com/problems/reverse-linked-list/description/)
### 문제 이해
linked list 의 head가 주어질때, list를 뒤집어서 반환하라

### 추상화
반복문을 이용한 방법
이전 포인터, 현재 포인터, 앞 포인터를 조작하여 순회하며 뒤집자.
현재 포인터의 next로 다음 노드를 가져오고
현재의 next를 이전 노드로
이전 포인터를 현재 노드로
현재 포인터를 다음 노드로
### 계획

### 코드
```python
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        cur = head
        while cur:
            nxt = cur.next
            cur.next = prev
            prev = cur
            cur = nxt
        return prev
```
재귀
```python
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        new_head = self.reverseList(head.next)
        head.next.next = head
        head.next = None
        return new_head
```
### 회고
머리아프다 linked list....
재귀 호출 코드는 직관적으로 이해가 잘 되지는 않는다