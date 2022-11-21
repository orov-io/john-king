up: build
	@echo "Booting john-king services with docker-compose"
	@docker-compose up -d

build:
	@echo "Building john-king services"
	@docker-compose build --parallel

logs:
	@echo "Tailing john-king containers logs"
	@docker-compose logs -f

down:
	@echo "Shutting down all john-king containers"
	-@docker-compose down

reboot: down up logs