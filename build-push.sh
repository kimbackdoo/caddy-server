docker buildx build --platform=linux/amd64,linux/arm64,linux/arm/v7 -t ho0on/caddy:latest caddy --push

docker buildx build --platform=linux/amd64,linux/arm64,linux/arm/v7 -t ho0on/tls-check:latest tls-check --push

docker buildx build --platform=linux/amd64,linux/arm64,linux/arm/v7 -t ho0on/web:latest web --push
