[link](https://leetcode.com/problems/sum-of-two-integers/description/)
### 문제 이해
+, - 연산자 없이 주어진 두 숫자의 합을 반환
-1000<=a,b<1000
### 추상화
비트 연산?
or 0 0 -> 0 / 1 1 -> 1 / 1 0 -> 1
and 0 0 -> 0 / 1 1 -> 1 / 1 0 -> 0

01 or 10 = 11
10 xor 11 = 01
10 and 11 = 10
A xor B or (A and B) left shift 1 

xor가 sum에 해당
and가 carry에 해당
sum과 carry를 구하고 carry shift 1해서 더해주면 합은 됨
a와 b를 sum과 carry로 업데이트
carry가 0이 될때까지 반복
### 계획
### 코드
```python
class Solution:
    def getSum(self, a: int, b: int) -> int:
        bitmask = 0xFFFF
        while b != 0 :
            s = (a ^ b) & bitmask
            c = ((a & b) << 1) & bitmask
            a = s
            b = c
        limit = 0x7FFF
        return a if a <= limit else ~(a ^ mask)
```
### 회고
파이썬은 자료형의 크기가 제한이 없기 때문에 음수 그냥 놔두면 carry 1이 안사라짐
	mask로 처리해서 범위 이상의 값은 버려줘야함