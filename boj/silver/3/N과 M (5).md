[link](https://www.acmicpc.net/problem/15654)
### 문제 이해
N개의 자연수와 자연수 M이 주어짐
조건을 만족하는 길이가 m인 수열
N개의 자연수는 모드 다른수
1 <= M <= N <= 8
### 추상화
순열 구현하는데 range를 이용하는게 아니라 입력받은 배열을 이용
DFS + 백트래킹

### 계획
N개의 자연수 입력받은 뒤 오름차순 정렬
재귀함수를 만들자.
	선택한 숫자를 저장하는 배열과 입력받은 배열을 클로저로 선언
		저장 배열은 M 크기로 초기화
	시작 idx과 depth를 파라미터로
선택 - 탐색 - 취소
	첫 인덱스부터 마지막 인덱스까지 순회하면서
	선택한 배열에 값을 추가하고 dfs 재귀함수 호출
		재귀호출시마다 idx+1, depth+1 하여 전달
	 depth 가 M이 되면 선택 저장 배열 출력
조합이 아니라 순열
	1 3 / 3 1 모두 출력해야함
	visited와 picked 분리
### 코드
```python
import sys

input = sys.stdin.readline


def solve(arr: list[int], N: int, M: int):
    arr.sort()
    visited = [False] * N
    picked = [0] * M

    def dfs(depth: int):
        if depth == M:
            print(" ".join(map(str, picked)))
            return
        for i in range(N):
            if visited[i]:
                continue
            picked[depth] = arr[i]
            visited[i] = True
            dfs(depth + 1)
            visited[i] = False

    dfs(0)


if __name__ == "__main__":
    N, M = map(int, input().split())
    arr = list(map(int, input().split()))
    solve(arr, N, M)

```
### 회고
지난번 문제의 depth와 picked 인덱스를 이용한 할당을 적용함
순열을 출력해야한다는 걸 문제 한참 풀다 파악함
	비슷한 문제라도 문제 파악을 좀더 꼼꼼히