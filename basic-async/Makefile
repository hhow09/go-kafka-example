up:
#	use timestamp as log file name
	LOG_FILE=$$(date +%s) docker compose up -d
up-rebuild:
	LOG_FILE=$$(date +%s) docker compose up -d --force-recreate --build consumer --build producer --build log-parser
down:
	docker-compose down --volumes