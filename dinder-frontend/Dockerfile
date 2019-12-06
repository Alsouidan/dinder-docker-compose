FROM node:12.2.0-alpine as build

WORKDIR /usr/src/app

COPY package.json yarn.lock ./

RUN yarn

COPY . ./

EXPOSE 3000

RUN yarn build

CMD ["npm","start"]