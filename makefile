registry:=ghcr.io
username:=fajar-islami
image:=github.com/alwisisva/twitter-app
tags:=latest


.PHONY: help
help: ## Show help command
	@printf "Makefile Command\n";
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: readenv
readenv: ## not work
	export $(cat .env | xargs -L 1)

.PHONY: migrate
migrate: ## Create Migrations file, example : make migrate name="xxxx"
	@if [ -z "${name}" ]; then \
		echo "Error: name is required \t example : make migrate name="name_file_migration";" \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations '${name}'


migrate.up: ## Up migration, example : make migrate.up envfile=.env.test
	go run cmd/migrate/main.go -envfile=${envfile}

migrate.rollback: ## Up rollback, example : make migrate.rollback -rollback=true envfile=.env.test
	go run cmd/migrate/main.go -rollback -envfile=${envfile}

.PHONY: migrate.version
migrate.version: ## Check out to history force version, example make migrate.force -ver=12341212121 envfile=.env.test
	go run cmd/migrate/main.go -version=${version} -envfile=${envfile}

.PHONY: migrate-steps
migrate.steps: ## teps looks at the currently active migration version.\n It will migrate up if n > 0, and down if n < 0.\n example : make migrate.steps steps=10  envfile=.env.test
	go run cmd/migrate/main.go -steps=${steps} -envfile=${envfile}


migrate.check: ## Check status migration, example : make migrate.check envfile=.env.test
	go run cmd/migrate/main.go -check -envfile=${envfile}


build:
	go build -o app cmd/app/main.go

run:
	./app

run.local:
	go run cmd/app/main.go

run.db:
	docker compose up psql_tweet_app -d

dc.check:
	 docker compose config
	 
dc.up: ## up compose image
	docker compose up -d

dc.logs: ## logs compose image
	docker compose logs -f

dc.stop: ## stop compose image
	docker compose stop

dc.down: ## rm compose image
	docker compose down -v


docker.build:
	docker build --rm -t ${registry}/${username}/${image}:${tags} .
	docker image prune --filter label=stage=dockerbuilder -f

docker.run:
	docker run --name ${image} -p 8080:8080 ${registry}/${username}/${image}:${tags}


docker.rm:
	docker rm ${registry}/${username}/${image}:${tags} -f
	docker rmi ${registry}/${username}/${image}:${tags}

docker.enter:
	docker exec -it ${image} bash

docker.enter-img:
	docker run -it --entrypoint sh  ${registry}/${username}/${image}:${tags}


push-image: dockerbuild
	docker push ${registry}/${username}/${image}:${tags}



enterpsql:
	docker exec -it psql_tweet_app psql -U ADMIN -d twitter-app


