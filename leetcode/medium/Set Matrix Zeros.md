[link](https://leetcode.com/problems/set-matrix-zeroes/)
### 문제 이해
MxN 배열 
원소에 0이 있으면 그 row와 col을 모두 0으로
M,N <= 200
in place로 해야함? 복사 없이 처리해야함
	반환하지 말고 입력으로 받은 배열을 직접 수정
### 추상화
그래프 순회
일반적으로 그래프를 순회하면서 0이 나온경우 상하좌우에 대해 0으로 설정하면서 탐색하면 배열 전체가 0이 됨
가장 단순한 방법
따라서 배열 전체를 순회하면서 0의 위치만 저장했다가 DFS/BFS로 0을 설정
-> O(MN)
어차피 해당 원소의 i row, j col 전체를 0으로 만드니까 0이 될 row와 col만 저장하면 O(M+N)으로 처리 가능
time : O(MN)
space : O(M+N)
### 계획
순회하면서 0 row, col 찾기
찾은 row, col에 대해 0으로 설정하기
### 코드
```python
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # time : O(MN)
        # space : O(M+N)
        zero_row = {}
        zero_col = {}
        M = len(matrix)
        N = len(matrix[0])
        for i in range(M):
            for j in range(N):
                if matrix[i][j] == 0:
                    zero_row[i] = True
                    zero_col[j] = True
        for row in zero_row:
            for j in range(N):
                matrix[row][j] = 0
        for col in zero_col:
            for i in range(M):
                matrix[i][col] = 0
```
상수 메모리
```python
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # time : O(MN)
        # space : O(1)
        M = len(matrix)
        N = len(matrix[0])
        first_row_has_zero = any(matrix[0][j] == 0 for j in range(N))
        first_col_has_zero = any(matrix[i][0] == 0 for i in range(M))
        for i in range(1,M):
            for j in range(1,N):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    matrix[0][j] = 0
        for i in range(1, M):
            if matrix[i][0] == 0:
                for j in range(N):
                    matrix[i][j] = 0
        for j in range(1, N):
            if matrix[0][j] == 0:
                for i in range(M):
                    matrix[i][j] = 0
        if first_row_has_zero:
            for j in range(N):
                matrix[0][j] = 0
        if first_col_has_zero:
            for i in range(M):
                matrix[i][0] = 0
```
### 회고
상수 메모리로 푸는 방법??
	기존 배열의 첫 원소에 플래그를 저장하면 됨
		원래 첫 row, col에 0이 포함된 경우는 별도의 bool에 저장
		첫번째 row, col을 제외한채로 순회를 돌면서 0으로 채우고
		두 bool에 따라 첫 row와 col에 대한 0 채우기 작업
	