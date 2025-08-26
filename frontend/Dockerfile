FROM node:22-alpine AS base
WORKDIR /app
COPY package.json .
RUN npm install
RUN npm i -g serve

FROM base as build
COPY . .
RUN npm run build

FROM build AS run
EXPOSE 3000
CMD [ "serve", "-s", "dist" ]