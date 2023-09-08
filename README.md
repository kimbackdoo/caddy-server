# Caddy Server Example

### 사용방법

- k8s에서 배포하기 위한 조건
    - Cloud 환경 또는 On-Premise 환경 구축되어 있어야함


```
# 로컬에서 실행
docker-compose up

# k8s 배포
./build-push.sh
kubectl create ns caddy
helm install caddy helm
```