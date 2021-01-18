.DEFAULT_GOAL := run

run:
	docker-compose up

stop:
	docker-compose down
	docker rmi grpc-services_app --force