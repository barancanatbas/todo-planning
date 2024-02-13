FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./

RUN npm install
RUN touch .env

RUN echo "VUE_APP_SERVICE_URL=http://127.0.0.1:8082" > .env

COPY . .

RUN npm run build

FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
