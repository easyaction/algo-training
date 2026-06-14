[link](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)
### 문제 이해
수익을 최대화 할 수 있는 매수/매도 시점 선택
매수 / 매도 시점은 달라야함
수익을 얻지 못하면 0 반환
### 추상화
투포인터?
l r 포인터 움직이면서 최소 / 최대 값을 찾기
일단 브루트 포스로 하면 O(N^2)
투포인터로 하면 O(N)?
3 1 5 0 3

### 계획
l, r = 0
r을 증가시키면서
	prices[r] < prices[l] -> L을 R로 초기화
	prices[r] > prices[l] -> profit update

테스트 케이스
[] -> 0
[1] -> 0
[ 1 2 ] -> 1
[ 1 1 ] - > 0
[ 2 1] - > 0
[1 2 3] - > 2
[ 2 1 3] - > 2
[3 1 2] - 1
[ 3 2 1 ] 0
3 1 5 0 3
### 코드
```python
```
### 회고
