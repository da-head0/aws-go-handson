# PHONY - 위조품이라는 뜻을 갖고 있음.
# 같은 이름과의 파일 충돌 회피 -> make build를 여러번 수행할 때, 동일한 이름의 파일이 존재해도 항상 실행되도록 함.
.PHONY: build clean deploy gomodgen

build: gomodgen
	# Go 모듈을 활성화하기 위한 환경 변수
	export GO111MODULE=on
	# go build로 hello, world 프로그램을 실행 가능한 바이너리로 컴파일, 결과물은 bin/ 에 저장됨
	# amd64 linux 운영체제에서 실행될 수 있도록 컴파일됨
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go

# bin/ 디렉토리 등 아래의 파일을 지움
clean:
	rm -rf ./bin ./vendor go.sum

# serverless 프레임워크를 실행하여 애플리케이션을 배포.
# --verbose : 배포 프로세스 중에 자세한 출력을 제공
deploy: clean build
	sls deploy --verbose

# gomod.sh 가 실행 가능한지 확인한 후 실행
# gomod.sh -> 지정된 모듈과 종속성을 기반으로 go.mod 파일을 실행하거나 업데이트
# 모듈 : Go에서 단일 단위로 함께 버전이 관리되는 관련 Go 패키지의 모음. go.mod 참고
gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
