BASE_NAME = cncamp
TAG = v1.0

# root:
# 	export ROOT=github.com/cncamp/golang

.PHONY: build_image
build_image:
	echo "building httpserver binary" && \
		docker build -t $(BASE_NAME)/httpserver:$(TAG) .
