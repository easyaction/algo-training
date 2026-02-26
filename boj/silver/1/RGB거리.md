[link](https://www.acmicpc.net/problem/1149)
### 문제 이해
N개의 집
거리를 선분으로 나타냄
1~n번째 집
r,g,b 셋 중한 색상으로 칠해야함
규칙 만족하면서 모든 집을 칠하는 비용 최솟값
인접한 집의 색상이 달라야함
N 주어지고, 각 집의 R,G,B 색상 칠하는데 드는 비용 이 n개의 줄에 입력됨
### 추상화
그래프 탐색?
단순하게 모든 경우의 수를 다센다고 하면 너무 숫자가 큼
DP?
i번째 집의 r, g, b 각각을 선택했을때의 최소값을 이전 i-1에서 가져올 수 있음
즉 cost_sum(i)(j) = cost(i)(j) + min(sum(i-1)(except j))
O(N)으로 계산 가능
### 계획
N을 입력받고, 2차원 배열로 코스트를 입력받는다.
min cost 배열을 Nx3 배열로 선언
비용 배열을 돌면서 값을 초기화
min_cost[0]은 수동으로 초기화
for i in range(N):
	if i == 0:
		min_cost[i] = cost[i]
	else:
		for j in range(3):
			min_cost(i)(j) = cost[i] + min(cost if idx != i else 1000 for idx, cost in enumerate(cost[i-1]))

### 코드
```python
import sys
input = sys.stdin.readline
MAX = 1e9

def solve(N:int):
    cost = [list(map(int,input().split())) for _ in range(N)]
    min_cost = [[MAX, MAX, MAX] for _ in range(N)]
    for i in range(N):
        if i == 0:
            min_cost[i] = cost[i]
        else:
            for j in range(3):
                min_cost[i][j] = cost[i][j] + min(cost if idx != j else MAX for idx, cost in enumerate(min_cost[i-1]))
    return min(min_cost[N-1])
    

if __name__ == "__main__":
    N = int(input())
    print(solve(N))
```
개선
```python
import sys
input = sys.stdin.readline

def solve(N:int):
    dp = list(map(int, input().split()))
    for _ in range(1, N):
        r, g, b = map(int, input().split())
        new_r = r + min(dp[1], dp[2])
        new_g = g + min(dp[0], dp[2])
        new_b = b + min(dp[0], dp[1])
        dp = [new_r, new_g, new_b]
    return min(dp)
    

if __name__ == "__main__":
    N = int(input())
    print(solve(N))

```
### 회고
기존 풀이는 공간 복잡도가 너무 큼.
	모든 상태를 저장할 필요가 없다
	어차피 N이 1000이라 실제 메모리 사용량은 그닥 차이나지 않기는 하지만, O(N)->O(1)로 개선 가능
	시간복잡도만 보지 말고 공간복잡도도 살펴보자.