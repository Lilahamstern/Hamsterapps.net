FROM node:14.1-alpine

# Override the base log level (info).
ENV NPM_CONFIG_LOGLEVEL warn

# Create app directory
RUN mkdir -p /usr/src/app
# Set work dir
WORKDIR /usr/src/app

# Copy package.json
COPY package.json /usr/src/app/
COPY tsconfig.json /usr/src/app/

# Install package.json
RUN npm install --silent

# Add mounts
ADD src /usr/src/app/src
ADD public /usr/src/app/public

CMD ["npm", "run", "start"]