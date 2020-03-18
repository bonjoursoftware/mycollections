FROM node:13.10.1-stretch-slim as ui-builder
WORKDIR /app
COPY package.json .
COPY package-lock.json .
COPY src ./src
RUN npm install && npm run build-dev