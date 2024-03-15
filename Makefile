include .env

.PHONY:
	default
	dev
	run
	build
	clean-processes
	clean-images
	clean-volumes
	nuke

default: run

deploy: build
	docker stack deploy --orchestrator=kubernetes -c docker-compose.yml cluster

dev:
	compose-watcher

run: build
	docker-compose up --force-recreate

build:
	docker-compose build

stop: 
	docker-compose down

clean-processes:
	docker rm -f $$(docker ps -qa)

clean-images:
	docker rmi -f $$(docker images -q)

clean-volumes:
	docker volume rm $$(docker volume ls -q)

nuke: stop clean-processes clean-images clean-volumes