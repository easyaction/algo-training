[link](https://www.hackerrank.com/challenges/mark-and-toys/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting)
### 문제 이해
장난감 가격 목록이 있고 예산이 있을때, 가장 살 수 있는 장난감 갯수?

### 추상화
그리디 + 정렬?
가격이 정렬되어있는 상태에서 가장 싼것부터 계산해나가면 될듯
정렬해놓고 누적합을 구하면서 초과하지 않는 마지막 인덱스 찾기

### 계획
우선 가격 배열을 정렬 O(NlogN)
루프 돌며 부분합을 계산 O(N)
k를 초과하면 index를 반환
TC
3 50
100 25 25 - k를 딱 만족하는 경우?
100 50 25 - 2번째부터 k 초과
1 2 3 - 전부 미달할때?
100 100 100 - 전부 초과할때?

### 회고
