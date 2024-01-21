FROM node:20 as runtime
WORKDIR /app

COPY . .

RUN npm install
RUN npm run build 

ENV HOST=0.0.0.0
ENV PUBLIC_URLAPI=https://url-go.lol

EXPOSE 4000
CMD node ./dist/server/entry.mjs