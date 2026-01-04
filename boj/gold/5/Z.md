[link](https://www.acmicpc.net/problem/1074)
### 문제 이해
크기가 2^n x 2^n 배열을 Z모양으로 탐색
0 1 4 5 
2 3 6 7
8 9 12 13
10 11 14 15
### 추상화
브루트포스부터 생각해보자.
0,1, 2, 3 방문부터 구현 이후 확장(재귀적으로?)
2^1 / 2(0,0) (0,1) (1,0) (1,1)
2^2 / 2(0,0), (0,2), (2,0), (2,2)
2^3 / 2 (0,0), (0,4), (4,0)( 4,4)
사분면을 결정하고, 앞 사분면까지의 숫자를 더해주기


### 계획
주어진 N에 대하여 2^(n-1) 을 기준으로 r, c가 큰지 작은지 비교하려 왼쪽오른쪽위아래를 판단

반환을 2^(n) * (0,1,2,3) + func(N-1, nxt_r,nxt_c)

nxt_r, nxt_c는 왼오위아래에 따라 2^(n-1) 을 뺀다
### 코드
```python
import sys

input = sys.stdin.readline


def solve(N: int, r: int, c: int) -> int:
    if N == 0:
        return 0
    half = 2 ** (N - 1)
    quad = (2 if r >= half else 0) + (1 if c >= half else 0)
    nxt_r = r - half if r >= half else r
    nxt_c = c - half if c >= half else c
    return half * half * quad + solve(N - 1, nxt_r, nxt_c)


if __name__ == "__main__":
    N, r, c = map(int, input().split())
    print(solve(N, r, c))

```
### 회고
갈피를 못잡아서 힌트 봤음
	문제를 정의하는 것 자체가 어려웠다.
	추후 다시 풀어볼 필요 있음
관련 키워드
	쿼드 트리, z-order
비트 연산으로도 풀어볼수 있다고 함
	