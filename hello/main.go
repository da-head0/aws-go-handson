// 패키지 선언 - 이 코드가 어떤 패키지에 속하는지 알려줍니다.
// Go 언어의 모든 코드는 반드시 패키지 선언으로 시작해야 합니다.
// main 패키지에 속한 코드임을 컴파일러에게 알려줍니다.
// main 패키지 - 프로그램 시작점인 main() 함수를 포함한 패키지.
// 프로그램이 실행되면, 운영체제는 프로그램을 메모리로 올리고(로드) 프로그램 시작점(main() 함수) 부터 한 줄씩 코드를 실행
// main() 함수를 포함한 패키지
package main

// 코드가 의존하는 외부 패키지를 가져옴
import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// 사용자 정의 type인 Response 생성. -> 코드를 더 읽기 쉽고, 이해하기 쉽게 만들음.
// APIGatewayProxyResponse - 요청이 API Gateway를 통해서 어떤 응답으로 리턴될건지를 정해줍니다.
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call (기본 Lambda 함수)
// context.Context 매개변수를 사용하고, 응답과 오류를 반환
func Handler(ctx context.Context) (Response, error) {
	// JSON으로 인코딩된 응답 본문을 저장하는 버퍼를 생성
	var buf bytes.Buffer

	// json.Marshal 함수로 JSON 응답 본문을 생성.
	// 본문은 message를 포함하는 map
	_, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	// json.Marshal 중에 오류가 있었는지 확인.
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	// json으로 인코딩된 body에서 문자를 escape 처리하여 json의 문자열의 형식이 적절하고 안전한지 확인.
	// json.Marshal 함수 자체가 필요에 따라 escape 문자를 포함하여 적절한 json 인코딩을 처리할 수 있으므로 생략 가능.
	// 위의 body -> _ 로 변경 필요
	//json.HTMLEscape(&buf, body)

	// StatusCode, 인코딩 정보, 응답 본문, Header를 사용하여 응답 구조를 초기화.
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
	// 준비된 응답, 성공적인 실행을 나타냄 (err 가 nil 이라는 값 반환)
	return resp, nil
}

// lambda 함수가 호출될 때 호출되는 함수
func main() {
	// Handler 함수를 진입점으로 사용하여 lambda 함수를 초기화함.
	lambda.Start(Handler)
}
