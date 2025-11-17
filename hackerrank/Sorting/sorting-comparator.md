[link](https://www.hackerrank.com/challenges/ctci-comparator-sorting/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting)
### 문제 이해
정렬시 사용하는 비교 연산자를 작성하라
Player라는 클래스를 만들고, decreasing order로 정렬되도록 작성
score 기준으로 정렬하되, 동일하면 name의 alphabet order로. 이것은 ascending 순서로
### 추상화
구현 문제?
score 기준으로 먼저 비교하고, 동일 스코어일때 알파벳 순서로 비교가 되도록.

### 계획
그대로 구현한다.
두 객체의 score를 먼저 비교하고, 그에따라 반환
동일 스코어일때는 name을 비교
string 비교가... 그냥 하면 되었었나?
### 회고
repr 객체를 문자열로 표현했을때 호출되는 함수?
init() constructor. 이때 self에 필드 등록해주면 되겠다.