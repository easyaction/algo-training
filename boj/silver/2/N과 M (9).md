[link](https://www.acmicpc.net/problem/15663)
### 문제 이해
N개의 자연수와 M이 주어질때, N개의 자연수중 M개를 고른 수열중 길이가 M인 수열을 모두 구하라
중복 출력 금지
### 추상화
DFS + 백트래킹 + 중복처리?
### 계획
입력받은 배열을 기준으로 수열을 DFS로 탐방하며 선택

선택한 숫자 목록을 해시맵에 저장하여 없는 경우에만 출력

### 코드
```python
import sys

input = sys.stdin.readline


def solve(N: int, M: int, arr: list[int]):
    ans = {}
    visited = [False] * N
    picked = [0] * M

    def dfs(depth: int):
        if depth == M:
            seq = tuple(picked)
            if not seq in ans:
                ans[seq] = True
            return
        for i in range(N):
            if not visited[i]:
                visited[i] = True
                picked[depth] = arr[i]
                dfs(depth + 1)
                visited[i] = False

    dfs(0)
    ans_list = ans.keys()
    for s in sorted(ans_list):
        print(" ".join(map(str, s)))


if __name__ == "__main__":
    N, M = map(int, input().split())
    arr = list(map(int, input().split()))
    solve(N, M, arr)

```
개선
```python
import sys

input = sys.stdin.readline


def solve(N: int, M: int, arr: list[int]):
    ans = []
    visited = [False] * N
    picked = [0] * M

    def dfs(depth: int):
        if depth == M:
            ans.append(" ".join(map(str, picked)))
            return
        prev = 0
        for i in range(N):
            if visited[i]:
                continue
            if prev == arr[i]:
                continue
            visited[i] = True
            picked[depth] = arr[i]
            prev = arr[i]
            dfs(depth + 1)
            visited[i] = False

    dfs(0)
    for s in ans:
        print(s)


if __name__ == "__main__":
    N, M = map(int, input().split())
    arr = list(map(int, input().split()))
    arr.sort()
    solve(N, M, arr)


```
### 회고
정렬을 먼저 처리하는게 효율적
	입력받은 숫자를 대상으로 정렬하는게 빠름
같은 depth내에서 이미 선택한 숫자를 스킵하도록 하여 중복제거
	중복제거를 당겨서 불필요한 연산을 줄이자