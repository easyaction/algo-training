[link](https://www.hackerrank.com/contests/software-engineer-prep-kit/challenges/task-scheduler-cooldown-multiple-machines/problem?isFullScreen=true)
### 문제 이해
m 개의 머신에서 task를 전부 처리할 수 있는 최소시간 구하기
tasks 배열에는 task type이 들어감
k는 cooldown 시간
	같은 task type을 k 동안 못함
### 추상화
cooldown을 최소화하면서 병렬처리 하는것이 포인트
m개만큼 잘라서 처리할 필요가 있다.
type마다의 갯수에 따라서 현재 처리할 최적의 작업을 선택해야할 필요가 있음
### 계획
Counter로 tasks 변환 -> O(N)
for loop을 돌며 step을 올린다 <- 모든 원소의 갯수가 0이 아닐때까지 순환
	가장 count가 많은 element 순으로 m이 찰때까지 뺀다
		cooldown에 포함되어있으면 제외
		뺀 원소는 cooldown 대상에 등록
		for loop 시작시 cooldown이 등록되어있으면 -1
			cooldown 목록과 값이 모두 등록되어야겠다..
포룹 종료시 step 반환
		
### 회고
