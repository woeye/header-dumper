binary := header-dumper
img := woeye/header-dumper
tag = $(shell cat .tag)

gen_tag:
	date +"%Y%m%d%H%M" > .tag

build:
	go build -o $(binary) -v

docker:
	docker build -t $(img):latest -t $(img):$(tag) .

push: docker
	docker tag $(img):$(tag) $(img):$(tag)
	docker push $(img):$(tag)

clean:
	rm $(binary)