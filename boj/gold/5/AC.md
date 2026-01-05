[link](https://www.acmicpc.net/problem/5430)
### 문제 이해
정수배열 연산
함수 R, D
R - 수의 순서 뒤집기
D - 첫번째 수 버리기. 배열 비어있는데 D 사용하면 에러
입력
Test case 수
함수(R,D 조합)
배열에 들어간 숫자 갯수
배열에 들어간 정수가 중괄호와 콤마 구분으로 입력됨
출력 - 남은 배열 출력
### 추상화
투포인터?
배열을 초기화 하고, L, R 포인터를 이용하여 삭제 연산 대신 포인터를 움직여 최종 배열 범위를 지정
R 연산이 들어올때마다 움직일 포인터를 변경 (L<->R)
L>R이 되면 에러

### 계획
배열 초기화
	중괄호 좌우를 떼고 콤마로 구분하기
L=0, R=len(arr)-1
함수에 따라 
	R- 뒤집어졌는지 판단하는 bool을 전환
	D- 현재 뒤집어졌는지 여부에 따라 움직일 포인터를 선택하여 인덱스 수정
		정방향 -> L 증가 역방향 -> R 감소
		D 연산 수행 전 L > R이면 error 출력
뒤집어졌는지 판단하는 bool에 따라 l,r 인덱스를 참고하여 순서대로 출력
	뒤집어진 경우 reversed()를 이용하여 거꾸로 출력 
### 코드
```python
import sys

input = sys.stdin.readline


def solve(funcs: str, arr: list[int]):
    l, r = 0, len(arr) - 1
    reversed_arr = False
    for func in funcs:
        if func == "R":
            reversed_arr = not reversed_arr
        else:
            if l > r:
                print("error")
                return
            if reversed_arr:
                r -= 1
            else:
                l += 1
    result = map(str, arr[l : r + 1] if not reversed_arr else reversed(arr[l : r + 1]))
    print("[{0}]".format(",".join(result)))


if __name__ == "__main__":
    T = int(input())
    for _ in range(T):
        [funcs] = input().split()
        N = int(input())
        [nums_str] = input().strip()
        arr = []
        if N > 0:
            arr = list(map(int, nums_str[1:-1].split(",")))
        solve(funcs, arr)

```
개선
```python
import sys

input = sys.stdin.readline


def solve(cmds: str, arr: list[str]) -> str:
    l, r = 0, len(arr) - 1
    is_reversed = False
    for cmd in cmds:
        if cmd == "R":
            is_reversed = not is_reversed
        else:
            if l > r:
                return "error"
            if is_reversed:
                r -= 1
            else:
                l += 1
    result = arr[l : r + 1]
    if is_reversed:
        result.reverse()
    return "[{0}]".format(",".join(result))


if __name__ == "__main__":
    T = int(input())
    for _ in range(T):
        cmds = input().strip()
        N = int(input())
        raw_input = input().strip()[1:-1]
        arr = raw_input.split(",") if N > 0 else []
        print(solve(cmds, arr))

```
### 회고
소소한 리팩토링
	입력 처리시 불필요한 언패킹 삭제 등등...
deque로도 풀 수 있겠다.
	l,r 대신 R이 들어왔을때 popleft() 혹은 pop()을 실행