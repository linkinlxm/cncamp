BASE_NAME = linkinlxm/cncamp
TAG = prometheus_v1.0

# root:
# 	export ROOT=github.com/cncamp/golang

.PHONY: build_image
build_image:
	echo "building httpserver prometheus binary" && \
		docker build -t $(BASE_NAME):$(TAG) .
		docker push $(BASE_NAME):$(TAG)

.PHONY: deploy
deploy:
	kubectl apply -f ./deployment-http-prometheus.yaml

.PHONY: replace
replace:
	kubectl replace -f ./deployment-http-prometheus.yaml