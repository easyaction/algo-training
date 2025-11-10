[link](https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays)
### 문제 이해
길이 n개 array
query a부터 b까지 k를 더함
index 시작이 1
모든 query를 돌리고 나서 max value 반환

### 추상화
diff array를 사용하여 부분합의 max를 구하는 것이 아닐까
### 계획
1. brute-force
직접 array를 전부 접근해서 값을 더해주면...
array 접근에 O(N)
query 처리에 O(N).. 
O(N^2)이 되면서 runtime error

2. 차분배열
array 접근에 O(1)
query 처리에 O(N)
부분합 구하는데 O(N)
O(N)으로 처리
### 회고
처음에 일부 타임아웃 발생
	n= 10^7 이라 이를 2회 이상 호출하면 타임아웃 발생
	굳이 부분합을 array로 구할필요 없이, 구한 차분배열에서 max sum만 구하면 된다.
차분 배열이라는 개념을 몰랐다면 못풀었을거같다.
	어떻게 O(N^2)를 줄일것인가를 생각해보자.
	부분합
		i번째 숫자의 합을 구한 배열을 선언
		An = Sn - S n-1
		을 이용한 트릭
	차분배열
		변경사항을 실제로 적용하지 말고, 변경사항에 대한 배열을 만들자
		실제 필요한 계산은 한번만 수행
