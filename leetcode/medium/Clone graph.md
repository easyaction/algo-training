[link](https://leetcode.com/problems/clone-graph/description/)
### 문제 이해
그래프를 복사(deep copy)
0 <= val <= 100
### 추상화
그래프 순회?
그래프를 탐색하면서 새로운 노드를 만들기?

### 계획
101개 노드를 담을 리스트를 선언
주어진 노드를 순회하면서 val idx에 값이 없으면 Node를 그 val로 만들고 인접한 이웃을 탐색
탐색할때 해당 값이 리스트에 있으면 이미 visited이므로 무시하고 아니면 queue에 넣기
bfs로 하자
### 코드
```python
from typing import Optional
from collections import deque

class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if node is None:
            return None
        visited = {}
        visited[node.val] = Node(node.val, [])
        q = deque([node])
        while q:
            cur = q.popleft()
            for nxt in cur.neighbors:
                if nxt.val not in visited:
                    visited[nxt.val] = Node(nxt.val, [])
                    q.append(nxt)
                visited[cur.val].neighbors.append(visited[nxt.val])
        return visited[node.val]
```
### 회고
dfs 재귀로 풀었다면?
	속도가 더 빨라질 수 있을지?