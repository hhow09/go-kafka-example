up:
#	use timestamp as log file name
	LOG_FILE=$$(date +%s) docker compose up -d --scale consumer=3
up-rebuild:
	LOG_FILE=$$(date +%s) docker compose up -d --scale consumer=3 --force-recreate --build consumer --build producer --build log-parser
down:
	docker-compose down --volumes