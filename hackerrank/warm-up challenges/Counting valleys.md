[link](https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)

### 문제 이해
해수면 기준 보다 높아졌다가 낮아짐 = 산
해수면 기준 낮아졌다 돌아옴 = 계곡
U = up D = down
지나간 계곡 갯수?

### 추상화
U = +1 D = -1
음수로 변한 횟수를 구하라.

### 계획
1. brute-force
리스트 순회
	U면 +1 D면 -1
	0보다 낮아지는 횟수를 셈
시간 O(N)
공간 O(1)

10^6 이니까 통과할 수 있지않을까?

### 회고
문제의 조건이 항상 sea level에서 시작해서 sea level로 끝난다고 했기때문에 0 => -1로 되는 갯수를 셌다.
하지만 개념적으로는 계곡이 끝나는 지점에서 갯수를 세는게 더 자연스럽겠다 싶음
	-1 =>0일때...