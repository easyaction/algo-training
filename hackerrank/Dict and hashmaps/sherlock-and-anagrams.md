[link](https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps)
### 문제 이해
주어진 문자열에 대해 substring중 anagram pair 갯수 구하기
substring에 중복이 허용되어야함
	예를들면 mom이라고 했을때, m, o, m, mo, om, mom 총 6가지 substring

### 추상화
문자열의 substring을 구하고, 이에 대한 문자 구성이 일치하는 경우를 구해야한다.
hashmap에서의 해시충돌을 어떻게 처리할 것인가?

### 계획
substring 배열 만들기
	이게 문제
	brute-force해보면
		length = len of substring=1->len(s)
			idx=0 -> len(s)-len of substring
				ss = s[idx:idx+length] 
Counter로 변환
	별도의 포인터로 표현되나?
		python이 포인터...?
	테스트해보면 될듯
구성요소 비교하여 동일한 경우 찾기?
### 회고
Counter 써서 해보려다가, 필요 없다는것을 깨달았음
	결국 substring의 문자 구성만 정렬해서 만들면 됨
		근데 정렬쓰니까 문자열 길이만큼 매번 호출하게됨..
		
