BUCKET_NAME=stivenramireza-lambda-secure-software
BUILD_NAME=authorizer.zip
build:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	build-lambda-zip -o $(BUILD_NAME) bootstrap

publish: build
	aws s3 cp $(BUILD_NAME) s3://$(BUCKET_NAME)/$(BUILD_NAME)