docker:
	make web -C web
	docker build . -t web:latest
	docker image tag web:latest localhost:5000/web:latest
	docker push localhost:5000/web:latest	
