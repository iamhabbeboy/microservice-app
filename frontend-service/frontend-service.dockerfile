FROM node:lts-alpine

# install simple http server for serving static content
RUN yarn global add http-server
RUN mkdir app/
# make the 'app' folder the current working directory

COPY . ./app
WORKDIR /app

# install project dependencies
RUN yarn install

RUN yarn build
# RUN ls

EXPOSE 8080
CMD [ "http-server", "dist/" ]