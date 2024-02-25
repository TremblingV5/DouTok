# Guidelines

> This file is used to tell everyone how to build, deploy and use DouTok.

## How to deploy the related dependencies

Run `docker-compose -f ./deploy/env.yml` to start all dependencies.

## How to build backend

For `User`, `UserDomain`, `Comment`, `CommentDomain`, `Favorite`:

1. Copy `./config/vscode_launch.jsonc` to `./vscode`
2. Use vscode to run these modules.

For others:

Run each modules by using `go run ./applications/xxx/`

## How to build frontend

## How to use DouTok
