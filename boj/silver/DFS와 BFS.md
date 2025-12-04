[link](https://www.acmicpc.net/problem/1260)
### 문제 이해
DFS/BFS 구현
입력 : 
노드 수 N, vertex 수 M, 시작할 노드 번호 V
m개의 줄에 걸처 연결된 두 노드 번호
그래프는 양방향
### 추상화
그래프를 구현하고, dfs/bfs를 각각 구현
그래프는 방문 여부를 저장하는 visited 배열과 간선 정보를 기록하는 2차원 배열 graph로 구성

### 계획
그래프 초기화
  visited 배열 선언(크기 N)
  간선 갯수만큼 돌면서 x y / y x = true 설정(양방향)
dfs / bfs 구현
dfs - 재귀 호출로 구현
bfs - deque 사용
### 코드
초기 형태
```python
import sys
sys.setrecursionlimit(100000)
from collections import deque


def solve(graph, V):
  def dfs(graph, visited, node):
    if visited[node]:
      return []
    visited[node] = True
    result = [node]
    for adj in sorted(list(graph[node].keys())):
      if not visited[adj]:
        result.extend(dfs(graph, visited, adj))
    return result

  def bfs(graph, visited, node):
    queue = deque([node])
    result = []
    while queue:
      n = queue.popleft()
      if not visited[n]:
        result.append(n)
        visited[n] = True
        for adj in sorted(list(graph[n].keys())):
          if not visited[adj]:
            queue.append(adj)
    return result
        
  print(" ".join(map(str, dfs(graph, [False] * (len(graph) + 1), V))))
  print(" ".join(map(str, bfs(graph, [False] * (len(graph) + 1), V))))


if __name__ == '__main__':
  [N, M, V] = input().strip().split()
  graph = [{} for _ in range(int(N)+1)]
  for _ in range(int(M)):
    [X,Y] =input().strip().split()
    graph[int(X)][int(Y)] = True
    graph[int(Y)][int(X)] = True
  solve(graph, int(V))
```
개선
```python
import sys
from collections import deque

sys.setrecursionlimit(100000)
input = sys.stdin.readline


def solve(graph, V, N):
    def dfs(node, visited):
        visited[node] = True
        result = [node]
        for nxt in graph[node]:
            if not visited[nxt]:
                result.extend(dfs(nxt, visited))
        return result

    def bfs(start, visited):
        q = deque([start])
        visited[start] = True
        result = []
        while q:
            cur = q.popleft()
            result.append(cur)
            for nxt in graph[cur]:
                if not visited[nxt]:
                    visited[nxt] = True
                    q.append(nxt)
        return result

    visited = [False] * (N + 1)
    dfs_order = dfs(V, visited)
    print(" ".join(map(str, dfs_order)))

    visited = [False] * (N + 1)
    bfs_order = bfs(V, visited)
    print(" ".join(map(str, bfs_order)))


if __name__ == "__main__":
    N, M, V = map(int, input().split())
    graph = [[] for _ in range(N + 1)]

    for _ in range(M):
        x, y = map(int, input().split())
        graph[x].append(y)
        graph[y].append(x)

    # 정점 번호가 작은 것부터 방문해야 하므로 미리 정렬
    for i in range(1, N + 1):
        graph[i].sort()

    solve(graph, V, N)

```
### 회고
dfs를 stack으로도 풀어보자
"".join() 사용법 숙지
자료 구조 선택 기준을 다시 생각해 볼 필요가 있음
	과하게 복잡한 방식으로 구현하는 경향이 있음