[link](https://www.hackerrank.com/challenges/repeated-string/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)
### 문제 이해
문자열 s 가 반복됨 길이 n 만큼
이때, a의 갯수 세기
### 추상화
주어진 s에서의 a 갯수 x (n // 문자열 길이) + (n % 문자열 길이) 에서의 a 갯수
### 계획
Counter를 이용하여 s에서 a 갯수 세기 = cnt
n% 문자열길이로 자른 문자열에서 a갯수 세기 = cnt_2
return cnt * n // 문자열 길이 + cnt_2
시뮬레이션 
s=abcac,  n=10
cnt = 2
cnt_2 = 0
n // 5 = 2
n % 5 = 0
cnt * 2 + 0 = 4
s=aba, n=10, len(s) = 3
cnt = 2, cnt_2 = 1, n // len(s) = 3, n % len(s) = 1
2 * 3 + 1 = 7

시간 복잡도 O(N)
	n <= 10 ^ 12 이지만, s가 <= 100이라 괜찮다.
	count를 s에 대해서만 하므로.

공간 복잡도 O(1)
### 회고
runtime error 발생
	아마 index 접근에 오류가 있었을 것..
	Counter(s).get("a")가 없을수도 있었음
		s에 a가 없는 케이스
의식적으로 test case 5개 만드는것 해야함
	엣지케이스, 빈값 만들기
문자열 전체를 count할 필요가 없다면, Counter 사용 과함
	.count()를 사용해도 됨.