[link](https://www.acmicpc.net/problem/5525)
### 문제 이해
N+1개의 I와 N개의 O로 번갈아가면서 이뤄진 문자열을 Pn이라고 함
P1 = IOI
P2 = IOIOI

I와 O로만 이뤄지진 문자열 S와 정수 N이 주어질때, S 안에 Pn이 몇군데 포함되어있는지 구하라
입력 - 
N
S의 길이 M
S
### 추상화
브루트포스(투포인터)
문자열을 구성하고 있는 Pn들을 구한다.
i가 나올때까지 l pointer 이동
l부터 r을 증가시키는데 o, i가 번갈아가면서 나오는지 확인
	두개씩 보면 되겠다.
oi가 번갈아가면서 나올때마다 counter를 올리고, 패턴이 깨지면 Pn을 저장
그런식으로 문자열 내에 Pn을 모두 구한다.
주어진 입력 n에 대하여 크거나 같은 Pn만 걸러서 아래를 계산한다.(계산 결과가 0 이상인 경우만 합산)
Pn - N + 1
포인터 두개를 쓸 필요는 없고, 하나로 이동시키면서 하면 됨
### 계획
idx = 0
pn_arr = []
while idx < len(S) :
	s[idx]가 I가 나올때까지 idx++
	도달하면, idx+2 < len(S)일때, 
	 idx+1, idx+2가 O, I면 cnt++
		 idx += 2
	아니면 idx += 1, cnt append
pn_arr에 대하여 Pn - N +1을 각각 구하고, 양수인 값만 필터하여 합산
	 
### 코드
```python
import sys

input = sys.stdin.readline


def solve(S: str, N: int) -> int:
    idx = 0
    pn_arr = []
    while idx < len(S):
        if S[idx] != "I":
            idx += 1
            continue
        cnt = 0
        while idx + 2 < len(S) and S[idx + 1] == "O" and S[idx + 2] == "I":
            cnt += 1
            idx += 2
        if cnt != 0:
            pn_arr.append(cnt)
        idx += 1
    sum = 0
    for pn in pn_arr:
        cnt = pn - N + 1
        if cnt > 0:
            sum += cnt
    return sum


if __name__ == "__main__":
    [N] = map(int, input().split())
    [M] = map(int, input().split())
    [S] = input().split()
    print(solve(S, N))
```

개선
```python
import sys

input = sys.stdin.readline


def solve(S: str, N: int) -> int:
    idx = 0
    result = 0
    while idx < len(S):
        if S[idx] != "I":
            idx += 1
            continue
        pn = 0
        while idx + 2 < len(S) and S[idx + 1] == "O" and S[idx + 2] == "I":
            pn += 1
            idx += 2
        if pn >= N:
            result += pn - N + 1
        idx += 1
    return result


if __name__ == "__main__":
    [N] = map(int, input().split())
    [_] = map(int, input().split())
    [S] = input().split()
    print(solve(S, N))

```
### 회고
Pn을 구해서 다 끝나고 다시 결과값을 계산했는데, 패턴이 끝나는 시점에 바로 계산 가능하다.
	배열 없애고 직접 합산
	