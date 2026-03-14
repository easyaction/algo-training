[link](https://leetcode.com/problems/two-sum/description/)
### 문제 이해
정수 배열 nums
target int를 만드는 두 숫자의 idx 찾기
배열 길이 <= 10^4
### 추상화
배열을 순회하며 특정 인덱스에서의 target과의 차이를 구하고
차이와 일치하는 숫자를 찾으면 현재의 인덱스와 반환
배열 한번 순회하면 됨
O(N)?
### 계획
diff dict를 만듦
nums를 순회하면서
	target과의 차이를 구함
	이게 dict에 있으면 현재의 idx와 dict의 인덱스 값을 같이 반환
	없으면 차이를 key로 현재의 인덱스를 value로 넣기
### 코드
```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        diff = {}

        for i, n in enumerate(nums):

            d = target - n

            if d in diff:

                return [diff[d], i]

            else:

                diff[n] = i
```
### 회고
