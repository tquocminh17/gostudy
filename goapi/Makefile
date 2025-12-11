up:
	docker compose up -d
provision:
	docker compose up -d postgres
	docker compose run server ./main migrate
	${MAKE} up
