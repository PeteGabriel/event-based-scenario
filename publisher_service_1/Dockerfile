FROM node:lts-stretch as builder

WORKDIR /publisher_svc_1

COPY package*.json ./

RUN npm install

COPY . .

CMD node server.js