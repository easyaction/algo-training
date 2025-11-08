https://www.hackerrank.com/contests/software-engineer-prep-kit/challenges/merge-and-sort-intervals/problem?isFullScreen=true

#### 문제 이해
[시작, 끝]  -> interval
여러 인터벌이 주어질때, 순서대로 정렬하고 겹치는 구간을 합쳐라

#### 추상화
구현 문제라 특별히 추상화 할게 없는듯?
설명에 어떤 구현을 해야하는지 써져있었음

#### 계획
시작 값을 이용하여 정렬
첫 interval부터 차례대로 interval을 확인하여 머지할건 하고 스킵할것은 스킵

#### 회고
파이썬 문법이 오랜만이라 많이 틀림
	정렬시 sorted를 사용했는데, 반환된 리스트를 assign하지 않아 정렬이 제대로 동작하지 않음
		list.sort(key=)를 사용하자.
	max()가 있는데 굳이 if else로 구현
		뭐 이건 동작은 하니까?
		
엣지케이스 검증 안함
	n이 0, 1일때 동작하는 것 검증 안함
바라보고 있는 마지막 interval을 별도의 변수로 선언했는데, 그러다보니 for loop 에서 마지막에 값이 누락되었다
확실히 계획을 구체화 하고 코딩하는 것이 필요하다.
	엣지케이스에 대한 검증 필수!!