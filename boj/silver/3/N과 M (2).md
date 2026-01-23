[link](https://www.acmicpc.net/problem/15650)
### 문제 이해
자연수 n과 m이 주어짐
조건을 만족하는 길이가 m인 수열
1부터 n까지 중복없이 m개 고른 수열 오름차순 출력
1 <= M <= N <= 8
### 추상화
조합을 구현
m개의 for loop를 만들어서 순차적으로 고르면 되는데...
	이것을 어떻게 일반화 할지? 재귀적으로 호출?
dfs?
	사용했음을 나타내는 배열을 이용하여 숫자를 스킵

### 계획
dfs 함수 선언
	함수 인자로 N과 visited, count를 전달
	visited가 false인 숫자들을 반환하면서 재귀적으로 호출
	반환하는 숫자의 visited를 업데이트하면서 같이 전달
	count가 1이면 값만 반환
N 값에 따라 visited 초기화
count로 M 전달
### 코드
```python
import sys
input = sys.stdin.readline

def solve(N:int, M:int):
    def dfs(start:int, picked:list[int]):
        if len(picked) == M:
            print(" ".join(map(str, picked)))
            return
        for i in range(start, N+1):
            picked.append(i)
            dfs(i+1, picked)
            picked.pop()
    dfs(1, [])

if __name__ == "__main__":
    N, M = map(int, input().split())
    solve(N, M)

```
### 회고
dfs 구현에 대한 감을 못잡아서 힌트얻음
	백트래킹 개념
		선택 -> 탐색 -> 취소
		3C2를 본다고 하면 1 2를 본다음 1 3 을 봐야함
		이때 1을 선택하고 다음 단계를 탐색
			2를 선택하고, 탐색이 종료됨
			이 선택을 취소해야함 -> pop()
			그 다음 3을 선택 -> 탐색 -> 취소
		1 선택을 취소
		2를 선택 ...
itertools에 combinations 함수 구현과의 비교?