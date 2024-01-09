# frontend/Dockerfile

FROM node:latest

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

CMD ["ng", "serve", "--host", "0.0.0.0"]
