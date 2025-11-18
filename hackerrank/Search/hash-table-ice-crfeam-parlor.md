[link](https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search)
### 문제 이해
여행 동안에 예산과 아이스크림 가격 리스트가 주어짐
예산에 맞는 두가지 아이스크림의 인덱스를 반환
유니크한 페어가 꼭 존재
### 추상화
리스트 내에서 두 숫자의 합이 예산이 되는 경우?
two sum과 유사한것 같은데?
### 계획
비용에서 예산을 차감하면서 hashmap에 abs(value)를 key로, index를 value로 저장
배열을 순회하면서  hashmap을 조회해서 있는 두 값을 찾음
그 인덱스 두개를 반환
O(N)?
### 회고
 two sum과 동일한 접근이라고 예상했고, 실제로도 맞았다.
디버깅하면서 의도한대로 동작이 잘 되는지 확인했고, test case로 엣지 케이스도 확인
	값이 동일한 두 값이 있는경우를 어떻게 처리할 것인지?
문제에서 답이 unique one이라고 해서 쉽게 풀 수 있었음.
	여러개라면.. 조금 복잡해지겠다
	cost가 동일한 경우가 여러개가 되면 가짓수가 많이 증가할 수 있음
dictionary 접근할때 get()을 꼭 쓰자..
two sum때처럼 for loop을 두번 돌 필요가 없다.
	map에 넣으면서 값을 조회하면 됨..
