[link](https://www.acmicpc.net/problem/9019)
### 문제 이해
계산기. 0이상 10000 미만의 십진수 저장할 수 있는 레지스터
레지스터에 저장된 n을 변환
네자리 수 d1, d2, d3, d4
D - n 두배. 결과값 9999보다 크면 2n mod 10000을 저장
S - n에서 1 뺀 결과. 0이면 9999가 됨
L - 자릿수 왼쪽 회전
R - 오른쪽 회전
두 숫자를 입력받았을때, 첫번째 숫자에서 두번쨰 숫자로 가는 최소한의 명령어 나열

### 추상화
구현 + BFS
변수 4개를 만들고, 네자리수로 변환 / 네자리수를 변수 네개로 분해하는 함수 만들기
D - 네자리수로 변환하고 두배 연산 후 mod 후 변환
S - 네자리수로 바꾸고 -1 / 0이면 9999 처리
L / R - 변수 네개를 이용해서 순서 바꾸기
### 계획
### 코드
```python
import sys
from collections import deque

input = sys.stdin.readline


def to_integer(d1: int, d2: int, d3: int, d4: int) -> int:
    return ((d1 * 10 + d2) * 10 + d3) * 10 + d4


def to_4digit(num: int) -> tuple[int, int, int, int]:
    return num // 1000 % 10, num // 100 % 10, num // 10 % 10, num % 10


def D(num: int) -> int:
    return (num * 2) % 10000


def S(num: int) -> int:
    return 9999 if num == 0 else num - 1


def L(num: int) -> int:
    d1, d2, d3, d4 = to_4digit(num)
    return to_integer(d2, d3, d4, d1)


def R(num: int) -> int:
    d1, d2, d3, d4 = to_4digit(num)
    return to_integer(d4, d1, d2, d3)


func_to_str = {D: "D", S: "S", L: "L", R: "R"}


def test():
    TC = [5000, 9999, 0, 100, 10]
    for t in TC:
        print("num: ", t)
        print("D: ", D(t))
        print("S: ", S(t))
        print("L: ", L(t))
        print("R: ", R(t))


def solve() -> str:
    funcs = [D, S, L, R]
    start, end = map(int, input().split())
    visited = [""] * 10000
    visited[start] = "1"
    q = deque([start])
    while q:
        cur = q.popleft()
        for func in funcs:
            nxt = func(cur)
            func_name = func_to_str[func]
            if nxt == end:
                return visited[cur] + func_name
            if len(visited[nxt]) == 0:
                visited[nxt] = visited[cur] + func_name
                q.append(nxt)
    return ""


if __name__ == "__main__":
    T = int(input())
    for _ in range(T):
        print("".join(solve()[1:]))

```
개선
```python
import sys
from collections import deque

input = sys.stdin.readline


def D(num: int) -> int:
    return (num * 2) % 10000


def S(num: int) -> int:
    return 9999 if num == 0 else num - 1


def L(num: int) -> int:
    return (num * 10 + num // 1000) % 10000


def R(num: int) -> int:
    return (num // 10 + num % 10 * 1000) % 10000


func_to_str = {D: "D", S: "S", L: "L", R: "R"}


def test():
    TC = [5000, 9999, 0, 100, 10]
    for t in TC:
        print("num: ", t)
        print("D: ", D(t))
        print("S: ", S(t))
        print("L: ", L(t))
        print("R: ", R(t))


def solve() -> str:
    funcs = [D, S, L, R]
    start, end = map(int, input().split())
    prev = [-1] * 10000
    how = [""] * 10000
    prev[start] = start
    q = deque([start])
    cur = 0
    while q:
        cur = q.popleft()
        if cur == end:
            break
        for func in funcs:
            nxt = func(cur)
            func_name = func_to_str[func]
            if prev[nxt] != -1:
                continue
            prev[nxt] = cur
            how[nxt] = func_name
            q.append(nxt)
    target = end
    result = []
    while target != start:
        result.append(how[target])
        cur = prev[target]
    result.reverse()
    return "".join(result)


if __name__ == "__main__":
    T = int(input())
    for _ in range(T):
        print("".join(solve()))

```
### 회고
변수 네개로 쪼개고 다시 합치고 하는 구현은 불필요
	그냥 L, R을 계산하는게 더 간단
결과를 더하면서 저장하지 말고, 이전 단계를 저장하고 어떻게 그 단계에 도달했는지 저장하도록 변경
	마지막에 역추적하면서 결과를 합치기
