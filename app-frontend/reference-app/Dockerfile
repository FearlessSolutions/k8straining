# base image
FROM node:12.2.0-alpine

# set working directory
WORKDIR /reference-app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /reference-app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /reference-app/package.json
COPY package-lock.json /reference-app/package-lock.json

RUN npm ci

# start app
CMD ["npm", "start"]
