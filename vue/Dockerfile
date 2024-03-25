FROM node:18.1.0

ADD . ./app

WORKDIR /app

RUN npm install

RUN npm run build

EXPOSE 8080

CMD ["npm", "run", "serve"]