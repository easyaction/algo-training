[link](https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms)
### 문제 이해
배열의 원소간 최소 절대값은 얼마인지?
n=10^5
### 추상화
정렬한다음 다음 원소와의 최소 diff를 구하면서 최소값을 찾기?
	실제로는 순차적으로 정렬한다음, 두 숫자간의 간격이 최소인 값을 구하는 문제
정렬에 O(NlogN)
탐색에 O(N)
### 계획
리스트를 정렬
원소 순회하면서 다음 원소와의 diff를 구하면서 최소값 기록
반환
### 회고
이게 왜 그리디?
