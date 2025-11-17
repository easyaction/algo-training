[link](https://www.hackerrank.com/challenges/ctci-making-anagrams/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings)
### 문제 이해
두 문자열이 서로 아나그램이 되게 하기 위해 삭제해야하는 문자의 갯수 구하기
문자열길이 1<= a, b <= 10^4
모두 소문자
### 추상화
두 문자열의 교집합을 찾고, 이에 해당하지 않는 문자의 갯수를 구하는 것

### 계획
각 문자열의 Counter 선언
이를 정렬(26개뿐이니 O(1))
두 Counter간 subtract하고, 0이 아닌것들의 합을 구함

Counter를 안한다면, freq list을 써볼 수 있겠다.
freq list를 사용하자.

freq list를 각각의 문자열에 대하여 생성
a => z까지 차이를 구해서 0이 아니면 절대값을 취해 모두 더하면 답

### 회고
Counter 안쓰고 freq list로 처리
억지로 zip을 한다음 또 map을 해서 했는데, 억지다..
	오히려 연산량 증가한다. 없애자
억지 안쓰고 for loop을 이용하여 적절하게 처리하도록 변경

Counter를 쓰면 union을 활용할 수 있구나
	이 경우 각각의 문자 갯수를 전부 더하고, union을 빼주는 방식
어차피 26개 밖에 안되어서 뭘 해도 쉽다

Counter를 쓰고, 차분을 구한다음 abs해서 전부 더해주면 됨
	subtract()는 Counter 내부 element를 제거하는 것이 아니라 실제로 count를 계산한다
		(b에만 있는 값이라면 0-b= -b가 되면서 합쳐짐)
		Counter("abc") - Counter("ab") = Counter("a") 가 됨
		그래서 a-b 와 b-a를 더하면 A합집합B - A교집합B가 됨
	elements()는 iterator라서 길이를 구하고 싶다면 list로 변환해줘야한다
	update()를 이용하여 별도 변수 안쓰고 변경사항을 그대로 반영할 수도 있다
	
freq list를 a,b 각각에 대하여 만들었는데, 사실 하나만 만들고 a는 더하고 b는 빼는 식으로 for loop을 줄일 수 있다.

map(lambda)를 써서 abs를 적용했지만, 그냥 list comprehension으로 표현 가능
	sum(map(lambda x:abs(x), some_list))
	sum(abs(x) for x in some_list)

아예 abs 함수만 전달하는 것도 가능
	sum(abs, some_list)
	하지만 가독성을 위하여 list comprehension이 바람직
