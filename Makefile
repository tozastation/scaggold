develop:
	docker build . -t sgoffold -f ./Dockerfile.test
	docker run --name sgoffold --rm -it sgoffold
clean:
	docker stop sgoffold
	docker rmi sgoffold
statik:
	statik -src=./templates