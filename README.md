### [Serverless Framework](https://www.serverless.com/)란? 
- Deliver software with radically less overhead.

#### 장점
- `serverless deploy`라는 명령어 한 줄로 코드를 클라우드 환경에 배포할 수 있음
- 특정 클라우드 업체에 종속되지 않음
- AWS의 경우 CloudFormation을 활용하여 코드형 선언적 인프라가 내장됨
- 자동으로 확장됨
- idle 상태일 때 요금이 발생하지 않음
- 코드와 인프라를 함께 배포하여 바로 사용할 수 있는 서버리스 애플리케이션 제공
- 간단한 구문으로 AWS lambda 함수 등을 안전하게 배포할 수 있음
- Go, Node.js, Python, Java, C#, Ruby, Swift, Kotlin, PHP, Scala 등을 지원
- 개발, 스테이지, 프로덕션 등 다중 환경 지원

### 왜 AWS와 함꼐?
- 프로비저닝된 동시성 기능 출시, 처음부터 웜 인스턴스를 구성할 수 있도록 함.
- 코드를 변경할 필요 없이 원하는 프로비저닝 인스턴스 수에 대한 값을 설정하면 AWS Lambda가 항상 해당 수량의 웜업된 인스턴스가 작업을 대기하도록 보장함

### 공식 문서 & 설치 방법
- https://www.serverless.com/framework/docs/getting-started


[npm](https://www.npmjs.com/package/npm)을 통한 `serverless 모듈 설치
```
npm install -g serverless
```

Getting Started
```
git clone 이 레포
cd 레포
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