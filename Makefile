b:
	docker-compose up -d

d:
	docker-compose -f docker-compose.yml down --remove-orphans

ps:
	docker-compose ps