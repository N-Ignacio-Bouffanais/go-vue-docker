# Backend - Go
FROM golang:latest AS backend

WORKDIR /app

COPY ./backend .

RUN go mod download

RUN go build -o main .

# Frontend - Vue.js
FROM node:latest AS frontend

WORKDIR /app

COPY ./frontend .

RUN npm install

RUN npm run build

# Combine backend and frontend into a single stage
FROM golang:latest

WORKDIR /app

COPY --from=backend /app/main .
COPY --from=frontend /app/dist ./frontend/dist

CMD ["./main"]
