[link](https://leetcode.com/problems/climbing-stairs/)
### 문제 이해
n을 만들기 위해 1 또는 2를 더할 수 있을때, 방법 수 구하기?

### 추상화
1까지 가는데 가짓수 1
2는 1 + 1, 2 두가지
f(n) = f(n-1)(1을 선택) + f(n-2)(2를 선택)
f(3) = f(2) +f(1) = 3

### 계획
1 과 2는 직접 초기화
그 이상에 대해서 순회하면서 값 채우기
n-1과 n-2를 담을 변수 두개를 갖고 계속 업데이트 해가며 값을 업데이트

### 코드
```python
class Solution:
	def climbStairs(self, n: int) -> int:
		if n <= 2:
			return n
	fn_1 = 2 # f(n-1), init f(2)
	fn_2 = 1 # f(n-2, init f(1)
	fn = 0
	for i in range(3, n+1): # f(3)부터 f(n)까지
		fn = fn_2 + fn_1
		fn_2 = fn_1
		fn_1 = fn
	return fn
```
### 회고
