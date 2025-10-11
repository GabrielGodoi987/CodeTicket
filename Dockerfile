FROM node:24-slim

RUN apt update && apt install -y openssl procps

RUN npm install -g @nestjs/cli@11.0.0

WORKDIR /home/node/app

USER node

CMD tail -f /dev/null