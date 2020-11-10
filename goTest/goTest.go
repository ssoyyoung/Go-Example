package main

func main() {
	// Go에서는 간편하게 사용할 수 있는 테스트 프레임워크를 내장하고 있다
	// go test 명령을 실행하여 테스트 코드를 실행할 수 있다.
	// go test는 현재 폴더에 있는 *_test.go 파일들을 모두 테스트 코드로 인식하고 일괄 실행한다.

	// 테스트 파일은 testing 이라는 표준 패키지를 사용한다.
	// 따라서 testing 패키지를 import 하고 테스트 매서드를 작성해야한다.
	// 테스트 매서드는 TestXxx와 같은 특별한 메서드명을 갖는데, 앞의 Test는 해당 메서드가 테스트 메서드임을 알리는 것이고 Xxx는 임의의 메서드명으로 처음 글자는 항상 대문자
	// 메서드의 ProtoType은 아래와 같이 testing.T 포인터를 하나 입력으로 받으며 출력은 없다. 테스트 에러를 표시하기 위해 testing.T 의 Error(), Fail() 등의 메서드들을 사용한다.

}
