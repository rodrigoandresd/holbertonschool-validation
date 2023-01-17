FROM golang:1.15.8-buster

WORKDIR /app

COPY . .

RUN apt-get update
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
RUN apt install nodejs
RUN npm install -g markdownlint-cli
RUN apt install sudo -y
RUN apt install xz-utils
RUN apt install make -y
RUN apt install zip -y
RUN npm install -g markdownlint-cli
RUN rm -rf /usr/local/go

