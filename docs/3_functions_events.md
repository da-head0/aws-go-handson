https://www.serverless.com/framework/docs/providers/aws/guide/intro

### Functions
- 여러분의 코드는 AWS Lambda functions으로 배포되고 실행됩니다.
- 각 function은 실행과 배포의 독립된 단위로, 마이크로서비스와 유사합니다.
- 사용 예제
```
- 데이터베이스에 사용자를 저장
- 데이터베이스에 있는 파일 처리
- 스케줄이 지정된 동작 실행
```
과 같은 하나의 job을 실행하기 위해 사용됩니다.

```
functions:          # Your "Functions"
  usersCreate:
    events:         # The "Events" that trigger this function
      - httpApi: 'POST /users/create'
  usersDelete:
    events:
      - httpApi: 'DELETE /users/delete'
```

### Events
functions은 events로 인해 실행됩니다. 

Events의 예시
- REST API
- S3 Bucket에 업로드된 파일(이미지 업로드)
- Cloudwatch 스케줄 (매 5분마다 실행)
- SNS 토픽의 메시지
- Cloudwatch 알림 
등이 있습니다.

Serverless 프레임워크를 사용하면 설정한 event에 필요한 인프라를 자동으로 생성하고, function이 해당 event를 수신하게 합니다.


### Services
- 프레임워크의 조직의 단위 입니다.한 애플리케이션 안에 여러 서비스가 있을 수 있지만 프로젝트 파일 같은 것이라 생각하시면 됩니다.
```
service: users
 
functions: # Your "Functions"
  usersCreate:
    events: # The "Events" that trigger this function
      - httpApi: 'POST /users/create'
  usersDelete:
    events:
      - httpApi: 'DELETE /users/delete'
```