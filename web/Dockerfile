FROM node:14-alpine as build

WORKDIR /app
COPY package*.json yarn.lock ./
RUN yarn install
COPY public public
COPY src src
RUN yarn build

FROM nginx:1.19-alpine
ENV GOTHERM_API http://host.docker.internal:8888
COPY nginx/templates /etc/nginx/templates/
COPY --from=build /app/build /usr/share/nginx/html
