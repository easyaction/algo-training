[link](https://www.hackerrank.com/challenges/minimum-swaps-2/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays)
### 문제 이해
정렬되지 않은, 중복 없는 자연수 배열 arr
두 숫자를 swap한다 했을때, 배열을 정렬할 수 있는 최소 스왑 횟수를 구하라
배열 길이를 n이라고 할때, 1~n까지 차있음
	숫자 자체만 놓고 보았을때 어느 순서에 와야하는지 알수있음

### 추상화
모든 경우의 수를 따져보기에는 array의 길이 n이 크다. 
	10^5 
정렬 기억이 안나지만, 몇가지 해보니 아래와 같은 규칙으로 동작하면 정렬이 이뤄짐
left right 인덱스가 있을때
왼쪽부터 본인 인덱스에 맞는 값을 스왑해나가면 된다?
	오른쪽 부터 채워도 값은 동일한듯..?

### 계획
idx map을 만들어 value와 idx 관계를 저장
오른쪽 포인터로 해당 위치에 맞는 값을 찾아 스왑
포인터 왼쪽 이동
반복

### 회고
idx_map에서의 값 스왑과 arr 값 스왑이 굉장히 헷갈려서 테스트 케이스에서 에러가 많이남
	일단 변수 선언을 안했더니 인덱스와 값이 헷갈렸다
		idx+1이 correct value인데...
	그래서 아예 debug 출력을 해서 문제 해결했다
		print()로 디버깅했는데... 이게 오히려 시간이 더 걸리나?
idx_map 만들필요 없이 그냥 arr.index() 함수를 사용하면 간단하게 풀수있다.?? No
	index(value, start, stop)
	stop index는 역시 미포함
	index 찾는 것이 O(N)이므로, 순회때마다 idx를 찾으면 O(N^2)가 되어 시간초과
idx_map을 swap할때 value값이 키라는 것이 헷갈리기 쉬웠다
	변수를 명시적으로 선언하는것이 가독성이 좋겠다.
	

swap()을 함수 짜서 했는데, pythonic하게는 아래와 같이 할수있음
```python
arr[i], arr[j] = arr[j], arr[i]
```
