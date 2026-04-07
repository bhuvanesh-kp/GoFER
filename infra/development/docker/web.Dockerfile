FROM node:20-alpine

RUN apk add --no-cache bash

WORKDIR /app

COPY web/package*.json ./

RUN npm install 

COPY web ./

RUN ls && pwd

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]