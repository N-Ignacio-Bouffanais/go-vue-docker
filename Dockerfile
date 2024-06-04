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

COPY backend/.env .env

RUN go build -o main .

# Establecer las variables de entorno
ENV BOT_TOKEN=${BOT_TOKEN}

# Combine backend and frontend into a single stage
FROM alpine

WORKDIR /app

COPY --from=backend /app/main .
COPY --from=backend /app/.env .env
COPY --from=frontend /app/dist ./frontend/dist

ENV ENV_FILE_PATH=/root/.env

EXPOSE 8083

CMD ["./main"]
