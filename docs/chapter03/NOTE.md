> [!NOTE]
> Go 에서 배열은 거의 사용되지 않는다.
> 
> - Go 에서 배열의 크기를 배열의 타입을 일부로 간주하는 제한이 있다.
>     - e.g. `[3]int` 와 `[4]int` 는 다른 타입이다.
> - 그런 타입간의 변환을 시도할 수도 업ㅈㅅ다.

---

> [!NOTE] Go 런타임
> 모든 고급 언어는 해당 언어로 작성된 프로그램을 실행할 수 있도록 하는 라이브러리 세트에 의존하는데, Go 도 예외는 아니다. Go 런타임은 메모리 할당, 가비지 컬렉션, 동시성 지원, 네트워킹 그리고 내장 타입과 함수 같은 서비스를 제공한다.
> 
> Go 런타임은 모든 Go 바이너리에 컴파일 되어 포함된다. 이는 개발 언어로 작성된 프로그램을 수행하기 위해 반드시 따로 설치를 해야 하는 가상머신을 이용하는 언어와는 차이가 있다. 바이너리에 런타임을 포함하는 것은 Go 프로그램의 배포를 쉽게 하고, 런타임과 프로그램 간의 호환성 이슈에 대해 걱정할 필요도 없다.
