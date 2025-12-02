[link](https://www.hackerrank.com/challenges/swap-nodes-algo/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search)
### 문제 이해
이진트리
in-order traversal 왼쪽 자식노드 - 루트 - 오른쪽 자식노드
스왑
왼쪽 자식트리와 오른쪽 자식 트리를 교환
k의 배수인 경우 subtree를 교체
n개의 노드가 있는 tree가 주어짐 root = 1
t swap 연산
바뀐 트리의 node idx를 출력
a-left b - right
-1은 null node
queries - k 목록

### 추상화
in-order traversal을 구현하고, k의 배수일때 좌우를 바꿔서 호출하도록 해야하겠다.
	재귀호출로 구현
### 계획
in-order traversal 구현
인자로 node, k, h를 받음
h가 k의 배수이면
	node.left, node.right = node.right, node.left
if node.left != -1
	in_order(node.left, k, h+1)
print(node.data)
if node.right == -1
	in_order(node.right, k, h+1)

 main 처리
 tree 구축
 n만큼 순회돌면서
	 a, b를 받으면 left, right 할당
	 factory같은게 있어야하나?
트리 순회
 
### 회고
일단 tree와 순회 구현이 부정확했음
사실상 트리 순회 구현문제였음
재귀로 풀었는데, 재귀 깊이 한계가 1000으로 막혀있어 죽었음
	sys.setrecursionlimit(10 ** 6)로 늘려줄 필요 있음
	