FROM node:20.13-alpine

WORKDIR /app

COPY package.json ./

RUN yarn install

COPY . .

RUN yarn build

EXPOSE 5001

CMD [ "yarn", "start" ]
