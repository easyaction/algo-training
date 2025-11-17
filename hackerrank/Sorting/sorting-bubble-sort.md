[link](https://www.hackerrank.com/challenges/ctci-bubble-sort/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting)
### 문제 이해
버블 정렬
이중 포룹으로 배열을 순회하면서 두 원소가 감소하는 방향일때 두 값을 스왑
배열의 원소는 오른쪽으로 갈수록 커져야 한다.(정렬됨)
O(N^2)
정렬하는데 실제로 수행한 스왑 횟수 출력
처음 / 마지막 원소 출력
### 추상화
n <= 600이므로, 360000 3e5..이므로 그냥 구현해도 동작은 하겠다.
그냥 구현하고 swap 일어나는 횟수를 출력?
### 계획
그냥 충실하게 bubble sort를 수행하고, swap이 일어난 횟수를 세어 출력한다.
### 회고
그냥 이렇게 푸는게 끝이네?

문자열 포맷팅 "문자열".format()
{0}, {1} -> 인덱스로 지정. format()의 파라미터의 인덱스를 의미
{name} -> 변수명으로 지정. format(name=some_variable) 이런식으로 지정