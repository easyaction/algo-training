[link](https://leetcode.com/problems/insert-interval/)
### 문제 이해
interval(i) = (start_i, end_i)
겹치지 않는 인터벌
start의 오름순으로 정렬됨
newInterval = (start, end)
새 intervals를 반환하라.
intervals.len <= 10^4

### 추상화
interval을 순회하면서 새 인터벌과 기존 인터벌간의 merging을 처리
브루트포스하게 먼저 해보자.

### 계획
1. new start, end < old start < old end	
	1. append(new + rest intervals)
2. old start/ end < new start/end
	1. append old start / end
3. new start <= old start,end <= new end
	1. update (new start, end) to new => skip
4. old start <= new start end <= old end
	1. update (old start, end) to new
5. new start <= old start <= new end <= old end
	1. update (new_start, old end) to new
6. old start <= new start <= old end <= new end
	1. update (old start, new end) to new

old 1,2 / new = 4,8 -> 2 [(1,2)] new 4,8
old 3,5 / new = 4,8 -> 6 [(1,2)] new 3, 8
old 6,7 / new = 3,8 ->  3 [(1,2)] new 3, 8
old 8, 10 / new 3,8 -> 5 [(1,2)] new 3, 10
12, 16 / 3, 10 -> 1 [(1,2), (3,10)(12,16) ]


### 코드
```python
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        ans = []
        if len(intervals) == 0:
            return [newInterval]
        new_start, new_end = newInterval
        for i, (start, end) in enumerate(intervals):
            if new_start == -1:
                ans.append((start,end))
            elif new_end < start:
                ans.append((new_start, new_end))
                ans.append((start,end))
                new_start = -1
            elif end < new_start:
                ans.append((start,end))
            else:
                new_start = min(new_start, start)
                new_end = max(new_end, end)
        if new_start != -1:
            ans.append((new_start,new_end))
        return ans
```
간략화
```python
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        from_new = [newInterval]
        before_new = []
        for start, end in intervals:
            if from_new[-1][1] < start:
                from_new.append([start,end])
            elif end < from_new[-1][0]:
                before_new.append([start,end])
            else:
                from_new[-1][0] = min(from_new[-1][0], start)
                from_new[-1][1] = max(from_new[-1][1], end)
        return before_new + from_new
```
### 회고
계획을 복잡하게 할 것 없이 나눠서 처리하는게 간단
newInterval을 기점으로 왼쪽 부분, 오른쪽 부분 나눠서 처리
	겹치는 부분이 있을 경우 오른쪽 부분의 마지막 원소를 newInterval로 보고 이를 계속 업데이트
	없는 경우 왼쪽 부분에 들어가야하는지 오른쪽 부분에 들어가야하는지 판단하여 넣기
	