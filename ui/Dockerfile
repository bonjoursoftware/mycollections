FROM node:13.10.1-stretch-slim as ui-builder
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY src ./src
RUN npm run build-prod