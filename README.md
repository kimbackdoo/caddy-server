# Caddy Server Example

### 사용방법

- 조건
    - domain 있어야 함
    - minikube로 환경 구축하려면 복잡
    - Cloud 환경 또는 현재 ip를 외부에서 들어올 수 있도록 포트포워딩 필요


```
# 로컬에서 실행
docker-compose up

# k8s 배포
./build-push.sh
kubectl create ns caddy
helm install caddy helm
```