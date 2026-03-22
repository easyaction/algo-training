[link](https://leetcode.com/problems/longest-substring-without-repeating-characters/)
### 문제 이해
문자 중복 없는 부분 문자열의 최대 길이?
문자열에는 영문자, 숫자, symbol, space 포함
### 추상화
two pointer? sliding window?
좌우 포인터를 조작해가며 문자열 전체를 한번 스캔해서 조건에 맞는 부분 문자열을 찾기

### 계획
문자열 해시맵에 저장
길이는 R과 L의 차이로 구한다.
	R - L + 1
발견한적 있다면
	l이 last_idx보다 작은 경우에 l을 last_idx + 1
dict last idx 업데이트
길이 = R - L + 1
최대 길이 업데이트
### 코드
```python
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last_idx = {}
        max_length = 0
        left = 0
        for right, c in enumerate(s):
            if c in last_idx and last_idx[c] >= left:
                left = last_idx[c] + 1
            last_idx[c] = right
            max_length = max(max_length, right - left +1)
        return max_length
```
### 회고
예전에는 문자의 수를 세서 중복이 안생길때까지 갯수를 차감하면서 left를 조작
