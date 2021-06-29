#Challenge Makefile

start:
	docker-compose up --build -d \
	&& sleep 15 \
	&& docker-compose exec integration-api sh -c './populate_db.sh'

check:
	docker-compose exec integration-api sh -c 'go test test/integration_api_test.go' \
	&& docker-compose exec matching-api sh -c 'go test test/matching_api_test.go'

#setup:
#if needed to setup the enviroment before starting it
