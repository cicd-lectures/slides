FROM node:18-alpine

LABEL Maintainers="Damien DUPORTAL<damien.duportal@gmail.com>, Julien LEVESY<jlevesy@gmail.com>"

# Install Global dependencies and gulp 4.x globally
RUN apk add --no-cache \
  curl \
  git \
  tini

RUN npm install --global npm npm-check-updates

# Install App's dependencies (dev and runtime)
COPY ./npm-packages /app/npm-packages
# By creating the symlink, the npm operation are kept at the root of /app
# but the operation can still be executed to the package*.json files without ENOENT error
RUN ln -s /app/npm-packages/package.json /app/package.json \
  && ln -s /app/npm-packages/package-lock.json /app/package-lock.json

WORKDIR /app
RUN npm install-clean
## Link some NPM commands installed as dependencies to be available within the PATH
# There muste be 1 and only 1 `npm link` for each command
RUN npm link gulp

COPY ./gulp/tasks /app/tasks
COPY ./gulp/gulpfile.js /app/gulpfile.js

VOLUME ["/app"]

# HTTP
EXPOSE 8000

ENTRYPOINT ["/sbin/tini","-g","gulp"]
CMD ["default"]
