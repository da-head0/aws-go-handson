### [Serverless Framework](https://www.serverless.com/)란? 
- Deliver software with radically less overhead.

#### 장점
- `serverless deploy`라는 명령어 한 줄로 코드를 클라우드 환경에 배포할 수 있음
- 특정 클라우드 업체에 종속되지 않음
- 애플리케이션이 돌아가는 인프라의 자동적인 확장과 축소가 가능
- 트래픽이 없어 코드가 작동하지 않을 때 요금이 발생하지 않음
- 코드와 인프라를 함께 배포하여 바로 사용할 수 있는 서버리스 애플리케이션 제공
- Go, Node.js, Python, Java, C#, Ruby, Swift, Kotlin, PHP, Scala 등을 지원
- 개발, 스테이지, 프로덕션 등 다중 환경 지원

### 왜 AWS와 함께?
- Serverless 프레임워크의 기본 클라우드 제공사
- AWS 환경에서 Serverless 프레임워크의 문서에 있는 모든 기능이 기본적으로 적용됨

### 공식 문서 & 설치 방법
- https://www.serverless.com/framework/docs/getting-started


[npm](https://www.npmjs.com/package/npm)을 통한 `serverless 모듈 설치
```
npm install -g serverless
```

[go](https://go.dev/doc/install) 설치
```
- OS에 맞는 Go 다운로드

```

Getting Started
```
git clone https://github.com/da-head0/aws-go-handson.git
cd aws-go-handson
```

AWS Credential 정보 입력
```
export ... 
```


배포

```
sls deploy
```

특정 함수만 빠르게 배포
```
sls deploy function --function hello

# 코드 배포가 안 될 때  
sls deploy function --function hello --force
```

Endpoint 호출
```
curl -XGET $endpoint
curl -XPOST $endpoint -H "Content-Type: text/plain" -d '$입력할 내용' 
```

제거
```
serverless remove
```