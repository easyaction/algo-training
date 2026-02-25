[link](https://www.acmicpc.net/problem/16953)
### 문제 이해
정수 a를 b로 만들때 필요한 연산 구하기
연산 2를 곱하거나, 맨 오른쪽에 1을 더하거나
최소 연산횟수 구한 값에 1을 더한 값을 출력
만들수 없으면 -1
1<= A < B <= 10^9
### 추상화
bfs
먼저 a에서 b를 구하는 경우를 생각하면
모든 경우의 수를 그래프 탐색으로 찾는것
bfs로 각 경우를 탐색하고, b를 초과하면 -1 찾으면 반환

그리디
거꾸로 b에서 a를 만들 수 있을지 확인해보면?
끝이 1이 나오면 1을 제거
그게 아니라면 2로 나누기
나누다가 1이 아닌 홀수가 나오면 이건 만들수 없는 것

### 계획
b를 mod 10하여 1이 나오면 / 10 아니면 나누기 2
a보다 작아지면 -1반환
a와 같아지면 반복한 횟수 반환
### 코드
그리디
```python
import sys
input = sys.stdin.readline

def solve(A:int, B:int):
	# greedy
    cnt = 1
    while(A < B):
        if B % 10 == 1:
            B = B // 10
        elif B % 2 == 1:
            return -1
        else:
            B = B // 2
        cnt += 1
    if A == B:
        return cnt
    else:
        return -1

if __name__ == "__main__":
    A, B = map(int, input().split())
    print(solve(A, B))

```
bfs
```python
import sys
input = sys.stdin.readline
from collections import deque

def solve(A:int, B:int):
    # bfs
    q = deque([(A,1)])
    while(q):
        nxt, depth = q.popleft()
        if nxt == B:
            return depth
        elif nxt < B:
            q.append((nxt * 10 + 1, depth+1))
            q.append((nxt * 2, depth+1))
    return -1
            

if __name__ == "__main__":
    A, B = map(int, input().split())
    print(solve(A, B))

```
### 회고
그리디
	2로 나눌때 1이 아닌 홀수가 나오는 경우를 놓쳤었음.
	모든 경우의 수를 잘 확인해야함
bfs
	값이 증가만 하므로 중복처리 안해도 되는데, 값이 줄어든다면 중복제거하자.
	