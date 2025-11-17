[link](https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting)
### 문제 이해
평소보다 더 많이 쓰면 고객에게 알림
고객의 특정 날짜동안의 소비량의 중앙값의 2배를 쓰면 문자 보냄
특정 날짜동안 소비량 안차면 안보냄
특정 날짜 구간 d 총 n일 문자 몇번 받는지 계산
중앙값 배열 길이가 짝수인 경우에는 두 원소 값의 평균?
### 추상화
d window만큼 보면서 중앙값을 구하고, 조건이 맞는지 확인하여 counter 증가
중앙값을 구하는 방법? 정렬해서 가운데 값 혹은 가운데 두 값의 평균
결국 정렬을 매 윈도우 만큼 새로 할 수 없으니, 효율적으로 정렬에 들어가는 비용을 최소화 해야한다.
이를 위해 삽입 정렬을 해야한다. 왜? 앞뒤 원소를 넣고 빼면서 기존 정렬을 활용할 수 있음


### 계획
### 회고
