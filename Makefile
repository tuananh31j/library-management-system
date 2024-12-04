start-db:
	docker start postgres-container || docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:17-alpine

stop-db:
	docker stop postgres-container

reset-db:
	docker stop postgres-container || true
	docker rm postgres-container || true
	docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=postgres -p 5432:5432 -d postgres:17-alpine