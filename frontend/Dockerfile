FROM node:18.17.1

WORKDIR /frontend

RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./

# If you patched any package, include patches before install too
# COPY patches patches

RUN pnpm install

# Bundle app source
COPY . .

EXPOSE 3000
CMD [ "pnpm", "dev" ]