FROM node:18-bullseye-slim as ui-builder
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY src ./src
RUN npm run build-prod
