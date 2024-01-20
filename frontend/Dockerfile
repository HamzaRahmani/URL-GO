FROM node:20 as runtime
WORKDIR /app


COPY . .

RUN npm install
RUN npm run build 

ENV HOST=0.0.0.0
ENV PUBLIC_URLAPI=http://localhost:5050

EXPOSE 4000
CMD node ./dist/server/entry.mjs