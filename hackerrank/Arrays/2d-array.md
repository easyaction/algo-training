[link](https://www.hackerrank.com/challenges/2d-array/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays)
### 문제 이해
6x6 배열
hourglass 패턴? 
abc
  d 
efg
이 패턴대로 더하는 것을 hourglass sum이라 함

6x6에 있는 16개의 hourglass sum중 최대값을 구하라
### 추상화
hourglass 패턴을 index로 표현하면?
```python
for i in range(0,4):
  for j in range(0,4):
	abc = ar[i][j:j+3]
	d = ar[i+1][j+1]
	efg = ar[i+2][j:j+3]
```
이를 더하면 sum
다 구해서 max 뽑으면 되겠다.
### 계획
brute-force
6x6 배열을 순회하며(i,j) i=0 => 3 j=0 => 3
	abc+d+efg를 구함
	list에 넣는다
max를 리턴
복잡도 : O(N)/ O(N) ?
	하지만 6x6으로 사이즈가 고정이기 때문에 문제는 없다

tc 
9로만 채우기
-9로만 채우기
0으로 채우기
9 -9 섞기

### 회고
if 문으로 max 업데이트 했는데, max()로 대체 가능
reduce 썼는데, sum()이 훨씬 빠르다. (C 내장 함수라서)
	커스텀 연산 아니면 쓰지말자.
	
```
sum(), len(), min(), max(), all(), any(), map(), filter(), sorted(), reversed(), ...
```