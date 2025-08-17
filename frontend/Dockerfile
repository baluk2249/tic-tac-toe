# Multi-stage Dockerfile for static frontend with nginx
FROM node:20 AS builder
WORKDIR /app
COPY . .
RUN npm install || true

# If you have a build step (e.g., React/Vue), add: RUN npm run build
# For plain HTML/JS/CSS, just copy files

FROM nginx:alpine
COPY --from=builder /app /usr/share/nginx/html
EXPOSE 80
