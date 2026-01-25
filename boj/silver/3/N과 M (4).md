[link](https://www.acmicpc.net/problem/15652)
### 문제 이해
자연수 n과 m이 주어짐
조건을 만족하는 길이가 m인 수열
같은 수 여러번 골라도 됨
비 내림차순
	수열 A가 A1 ≤ A2 ≤ ... ≤ AK-1 ≤ AK를 만족
1 <= M <= N <= 8
### 추상화
조합 구현 - 백트래킹(DFS)
중복 허용

### 계획
재귀함수를 만들자.
	시작 idx와 선택한 배열을 넘겨받는 함수
선택 - 탐색 - 취소
	다음으로 넘겨주는 숫자를 지금 선택한 숫자를 넘겨 그 숫자부터 순회하도록

### 코드
```python
import sys

input = sys.stdin.readline


def solve(N: int, M: int):
    def dfs(start: int, picked: list[int]):
        if len(picked) == M:
            print(" ".join(map(str, picked)))
            return
        for i in range(start, N + 1):
            picked.append(i)
            dfs(i, picked)
            picked.pop()

    dfs(1, [])


if __name__ == "__main__":
    N, M = map(int, input().split())
    solve(N, M)
```
개선
```python
import sys

input = sys.stdin.readline


def solve(N: int, M: int):
	picked = []
    def dfs(start: int):
        if len(picked) == M:
            print(" ".join(map(str, picked)))
            return
        for i in range(start, N + 1):
            picked.append(i)
            dfs(i)
            picked.pop()
    dfs(1)


if __name__ == "__main__":
    N, M = map(int, input().split())
    solve(N, M)
```
### 회고
클로저를 이용하자
	굳이 매번 picked를 전달할 필요 없음
picked list에 매번 원소를 append/pop 하는 대신 미리 길이 M의 배열을 선언하고 값을 업데이트하는 방법도 있음
	이때는 depth를 전달하여 종료조건 설정
	n,m이 작아서 큰 성능 차이는 없을듯