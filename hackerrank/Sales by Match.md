https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
### 이해
색상으로 짝지어진 양말 더미
숫자로 각 양말의 색상을 포현
몇개의 쌍이 맞는지 계산

#### 추상화
정수 배열에서 숫자 2개 될때마다 count +1

#### 계획
배열을 순회하며 각 숫자의 갯수를 계산
	hashmap에 k:숫자 v:갯수 를 저장
		Counter()를 써볼 수 있을까?
		우선은 list
각 갯수를 2로 나눈 몫을 합하여 반환
시간 복잡도 O(N)
공간 복잡도 O(N)?
n이 100미만

TC
7
1 1 1 1 1 1 1
10
1 2 3 4 5 6 7 8 9 10
9
10 20 20 10 10 30 50 10 20

#### 회고
Python은 나머지 연산자 쓰면 실수형 반환하므로, int로 변환하던, // 연산자를 사용해야함.
Counter를 사용해본다면?
	가능.
	hackerrank에서는 autocomplete가 되네?
Counter 적용하면서 map, reduce 오랜만에 써봄
	map(_function_, _iterable_, _/_, _*iterables_, _strict=False_)
	functools.reduce(_function_, _iterable_, _/_[, _initial_])
 아래와 같이 간소화

```python
def sockMerchant(n, ar):
	cnt = Counter(ar)
	result = reduce(lambda x, y: x+y, map(lambda x: x//2, cnt.values()))
	return result
```






