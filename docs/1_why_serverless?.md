[![](https://mermaid.ink/img/pako:eNpNUc1Kw0AQfpVlzrFmt0mb5uzVU2-SSzBbDTRNSVNRS8GfCAUDKjYomopCpQgFqxRRqk-0k3dwm9LqnnZnvr-Z7cC273AwoelyYjWIPBaIpxRnY4JJD28uMUrFKCZiMskuxlkyJWIaWUBMQrUlfg55j_7wBH_64joVz98ER8f4eEXEVywL6wuRnF1UV24XEc6GBAdT8fqJoyPpThaKUk1K3WZn8bwZ9QieDjB6y_mMLvlVHuzxoM5bLZL1I_HYw4cI7-PsWE5wMsa7l0Vatkq7CnfoNjH9zv2SBD8SOd08vXiIxfmQ_AvLQAGPB57tOnJVnbmSBeEu97gFprw6vGa366EFXQm026FfPWhsgxkGba5Au-nYId9w7Z3A9sCs2fWWrHLHDf1gc7H8_A8UaNqNLd_3lkT5BLMD-2CuaWWjYFRUtWTQUlljmlpU4ABMSvWCVqGMFTVaYlTTWVeBw1yCFkoVnTGtTHVDL6tF1ej-AsDn02o?type=png)](https://mermaid.live/edit#pako:eNpNUc1Kw0AQfpVlzrFmt0mb5uzVU2-SSzBbDTRNSVNRS8GfCAUDKjYomopCpQgFqxRRqk-0k3dwm9LqnnZnvr-Z7cC273AwoelyYjWIPBaIpxRnY4JJD28uMUrFKCZiMskuxlkyJWIaWUBMQrUlfg55j_7wBH_64joVz98ER8f4eEXEVywL6wuRnF1UV24XEc6GBAdT8fqJoyPpThaKUk1K3WZn8bwZ9QieDjB6y_mMLvlVHuzxoM5bLZL1I_HYw4cI7-PsWE5wMsa7l0Vatkq7CnfoNjH9zv2SBD8SOd08vXiIxfmQ_AvLQAGPB57tOnJVnbmSBeEu97gFprw6vGa366EFXQm026FfPWhsgxkGba5Au-nYId9w7Z3A9sCs2fWWrHLHDf1gc7H8_A8UaNqNLd_3lkT5BLMD-2CuaWWjYFRUtWTQUlljmlpU4ABMSvWCVqGMFTVaYlTTWVeBw1yCFkoVnTGtTHVDL6tF1ej-AsDn02o)

로컬 안에서만 배포해 봄
- 다른 사람도 내 서비스를 사용하게 해 보고 싶다면?

서버 안에서 코드 직접 빌드/배포
- State 가 있음.
- 서버가 갑자기 삭제되면? 복구하는 데 걸리는 시간은?
- Docker image등으로 언제든 다시 똑같은 상태의 서버가 생성되는 게 이상적.

도커 이미지로 서버에 컨테이너 생성
- 계속 동작하는 서버가 필요한 경우엔 이상적
- EC2의 경우엔 리소스를 다 사용하지 못할 수도
- 호출이 하루에 1번 온다면?

코드를 zip으로 압축해서 람다 배포
- serverless지만, AWS Console에서 클릭을 몇 번씩 수동으로 해줘야 함.
- 매번 zip 파일을 만들어 일일이 올리는 것이 번거로울 수 있음

Serverless 프레임워크 사용
- [README.md](../README.md) 문서 참조.