[link](https://www.hackerrank.com/challenges/ctci-ransom-note/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps)
### 문제 이해
잡지의 단어 목록이 주어지고, ransom note의 단어가 주어짐
왼쪽으로 오른쪽을 구성할 수 있는지 여부 판단
대소문자 구분함
잡지 단어 오려붙이는거니까 갯수도 맞아야할 것 같음
### 추상화
단어 목록이 주어질때, 단어 갯수 세고 이것이 동일한지 확인하는 문제라고 할수있다
단순하게는 hash map에 단어 목록을 구성하고, 주어진 단어가 hash map에 존재하는지 확인하면 되겠다.

### 계획
brute-force
주어진 magazine[m]을 이용하여 hashmap 생성
	key 단어 value 단어 갯수
note[n]을 순회하며 hashmap에 단어 있는지 확인
	갯수 -1
	갯수 모자라거나, 없으면 No 반환
Yes 반환

using Counter()?
hashmap을 Counter로 만들 수 있음

TC
6 3 - ok
Abc def ghi jkl mLo pqr
Abc def mLo
6 3 - fail(단어 갯수 불일치)
Abc def ghi jkl mLo pqr
Abc Abc mLo
6 3 - fail(단어 없음)
Abc def ghi jkl mLo pqr
Abc abc mLo
### 회고
원하는 output이 return이 아니고 print()였음
	확인하자...
dict 값 가져올떄 get() 사용하기
	값 없으면 None 반환
