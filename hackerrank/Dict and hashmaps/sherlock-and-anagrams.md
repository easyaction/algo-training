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
		정렬 안쓰는 방법?
			tuple을 써서 26개 문자에 대해서 갯수를 센다?
				Counter 쓰는거랑 뭐가 다르지..?
				Counter도 정렬이 필요해서 동일하지만, 튜플 선언하면 아님
				26개 문자 갯수를 리스트에 선언하고, tuple을 만들어서 동일
					hashable한게 tuple이라 이걸로 변환
					frozenset도 되는데, tuple이 만만하다
range의 index가 좀 헷갈린다
	stop은 미포함..
	계획 세울때 표기를 좀더 명확하게 해보자. 
		계획은 잘 썼는데, 가독성이 떨어져서 이게 맞는지 틀린지 알기 어려웠다.
시간 복잡도 O(N^3) -> O(N^2)으로 개선 가능
	substring을 매번 만들어서 O(L)이 소요됨
	substring을 따로 만들지말고, a, ab, abc 이런식으로 i는 두고 j를 확장하는 식으로 하면 전단계에서의 갯수에서 새로운 인덱스의 갯수만 더해서 처리가 가능
		슬라이딩 윈도우 관점으로 볼수 있겠다..