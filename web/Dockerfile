# build stage
FROM node:lts-alpine as build

WORKDIR /app

COPY package*.json ./
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN apk add make
RUN apk add build-base
# Install pnpm
RUN npm install -g pnpm

RUN pnpm install
COPY . .
RUN pnpm run generate

# production stage
FROM nginx:stable-alpine as production
COPY --from=build /app/dist /usr/share/nginx/html

# Copy the nginx configuration file
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
