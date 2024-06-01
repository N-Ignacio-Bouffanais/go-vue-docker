# Frontend - Vue.js
FROM node:latest AS frontend

WORKDIR /app

COPY ./frontend .

RUN npm install

RUN npm run build

# Backend - Go
FROM golang:alpine AS backend

WORKDIR /app

COPY ./backend .

RUN go mod download

RUN go build -o main .

# Combine backend and frontend into a single stage
FROM alpine

WORKDIR /app

COPY --from=backend /app/main .
COPY --from=frontend /app/dist ./frontend/dist

EXPOSE 8083

CMD ["./main"]
