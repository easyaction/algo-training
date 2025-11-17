[link](https://www.hackerrank.com/challenges/angry-children/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=greedy-algorithms)
### 문제 이해
unfairness = max(arr')-min(arr')
arr' - arr중 원소 k개 뽑은것

### 추상화
max - min이 최소가 되려면 숫자 구간이 최소여야한다
배열을 정렬하면 오른쪽 원소가 max 왼쪽 원소가 min
k 간격 만큼 떨어진 원소간의 차이를 계산해서 최소값을 찾으면 됨
### 계획
정렬 - O(NlogN)
i = 0 => len(arr) - k까지 range(len(arr)- k+1) - O(N)
arr[i+k-1] - arr[i]  min찾기
### 회고
계획할때 index를 정확하게 쓴게 도움이 되었음