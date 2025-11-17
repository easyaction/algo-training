[link](https://www.hackerrank.com/challenges/special-palindrome-again/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings)
### 문제 이해
특별한 문자열 조건
	모든 문자가 동일할때
	가운데 문자만 제외하고 모두 동일할때
어떤 문자열이 주어질때, 특별한 substring이 몇개 만들어질 수 있는지 계산하라
1 <= n <= 10^6
substring의 중복을 허용
영어 소문자
### 추상화
substring을 만들면서 그것이 조건을 만족하는지 체크
	이를 효율적으로 해야한다. 
		n=10^6 이므로...
회문 찾기와 유사하게 특정 문자를 기준으로 좌우로 window를 확장할 수 있겠다.
	회문의 부분집합이겠다..
substring의 길이에 따라 일반화를 해보자.
	1인 경우 모두 해당
		계산 스킵
	2인 경우
		문자 동일한 경우를 찾는다 O(N)
	3인 경우
		문자 동일한 경우 찾기 => 아래 조건의 부분집합
		좌우 문자 동일한 경우 찾기 O(N)
	4인 경우
		2인 경우에서 기존 문자와 좌우 문자 동일한 경우 찾기
	5인 경우
		3인 경우에서 좌우 문자 동일한 경우 찾기
	...
### 계획
brute-force
모든 substring에 대하여 조건 확인
	조건에 해당하는 경우 cnt++
i = 0 => len(s)-1
 j = i => len(s)-1
 O(N^2) => 시간 초과 뜰것
O(N log N) 내에 풀어야함

DP?
기존 계산 결과를 활용하는 방안
	substring 길이 1인 경우 스킵(문자열 길이 더하기)
	substring 길이 2인 경우 쭉 계산하면서 스택에 해당하는 케이스 넣기
스택을 이용하여 DFS해볼까..? 큐를 이용한 BFS?
가짓수가 많으니 dfs부터 해보자

dp에 저장하는 것은 ord(c)로 하자.
	해당 문자와 좌우가 동일하면 special
dp n x n 행렬 초기화... 미리 해놓을 필요가 있나
	이전 단계만 있으면 되니까 그때 그때 만들자
dp n 행렬 초기화

지금 DP와 DFS가 섞임. DP만 이용해보자.
DP...안쓰고 DFS사용

	
### 회고
문제 풀이 방향성을 제대로 잡지 못했다
	DP인지 DFS인지...
	DP로 풀면 시간 초과 날것같아서 중간에 선회
DFS로 풀었는데, 