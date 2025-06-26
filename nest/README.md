<p align="center">
  <a href="http://nestjs.com/" target="blank"><img src="https://nestjs.com/img/logo-small.svg" width="120" alt="Nest Logo" /></a>
</p>

[circleci-image]: https://img.shields.io/circleci/build/github/nestjs/nest/master?token=abc123def456
[circleci-url]: https://circleci.com/gh/nestjs/nest

<p align="center">A progressive <a href="http://nodejs.org" target="_blank">Node.js</a> framework for building efficient and scalable server-side applications.</p>

## Project setup

```bash
npm install
```

## Compile and run the project

```bash
# development
npm run start

# watch mode
npm run start:dev

# production mode
npm run start:prod
```

## Run tests

```bash
# unit tests
npm run test

# e2e tests
npm run test:e2e

# test coverage
npm run test:cov
```

## Linting & Formatting
```bash
# lint and autofix with eslint
npm run lint

# format with prettier
npm run format
```

## Nest CLI
```bash
# create a resouse
nest g resource [name]

# create a controller
nest g controller [name]

# create a service
nest g service [name]
```

## Init Modules 
```bash
# --- Generate Feature Modules ---

# Auth module for handling login, registration, and JWTs
nest g module modules/auth
nest g controller modules/auth
nest g service modules/auth

# Users module for user CRUD and profile management
nest g module modules/users
nest g controller modules/users
nest g service modules/users

# Posts module for post CRUD and interactions like likes
nest g module modules/posts
nest g controller modules/posts
nest g service modules/posts


# --- Generate Supporting Infrastructure Modules ---

# Database module for TypeORM configuration and migrations
nest g module modules/database

# Storage module for abstracting file uploads (R2/Azure)
nest g module modules/storage
nest g service modules/storage

# Cache module for managing Redis connections and caching logic
nest g module modules/cache
nest g service modules/cache
```

## Database

### Migrations
```bash
# don't forget to update the YOUR-MIGRATION-NAME part in the command
npm run typeorm migration:generate src/migrations/<YOUR-MIGRATION-NAME> -- -d src/modules/database/data-source.ts
```

### Seeding
```bash
npm run seed
```




