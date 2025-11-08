#### 문제 이해
두 문자열에 대하여 anagram인지 판단
#### 추상화
각 문자열 문자의 갯수를 세서 각각 문자 갯수가 동일하면 anagram이라고 판단할 수 있다.

#### 계획
문자열 순회
	char별로 hashmap에 갯수 추가
두번째 문자열때에 역으로 하나씩 갯수 감소
모든 문자가 0개 갯수면 ok
아니면 false
O(N)
#### 회고
python dict 순회시 items() method 이용
collections.Counter를 활용할 수도 있다고 함.
	알아서 문자별로 갯수 세고, 객체간에 + - 연산도 지원함
		-연산한다음 안비어있으면 False, 비어있으면 True 하면 되겠다.