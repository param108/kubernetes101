build:
	make web -C web

dockerv1: build
	make webv1 -C web
	docker build . -t web:v1

dockerv2: build
	make webv2 -C web
	docker build . -t web:v2
