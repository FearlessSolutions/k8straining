run_app:
	docker-compose up app

run_app_dev:
	docker-compose up storybook

down:
	docker-compose down

test_jest: docker-compose up jest

test_cypress: docker-compose up cypress

test_app:
	test_jest
	test_cypress

