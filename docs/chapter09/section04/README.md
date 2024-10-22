---
title: 9.4 모듈 관련 작업
aliases:
  - Working with Modules
---

다른 모듈들과 모듈 내에 패키지를 통합하는 방법을 살펴보자. 그 후에는, 직접 만든 모듈의 버전관리 및 배포와 pkg.go.dev, 프록시 모듈, 집계 데이터베이터와 같은 Go 의 중앙 집중적 서비스를 알아보자.

### [9.4.1 서드-파티 코드 가져오기](9.4.1.md)

---

### 9.4.2 버전 작업

기본적으로 Go 는 모듈을 프로젝트에 추가하면 의존성의 최신버전을 선택한다.

`go list` 명령어와 `-m` 인자를 사용하여 모듈의 이용가능한 버전을 확인할 수 있다.

```shell
go list -m -versions github.com/learning-go-book/simpletax
```

`go get` 명령어와 `@` 를 사용하여 특정 버전을 설치할 수 있다.

```shell
go get github.com/learning-go-book/simpletax@v1.0.0
```

---

### 9.4.3 최소 버전 선택

- TODO

---

### 9.4.4 호환되는 버전으로 업데이트

- TODO

---

### 9.4.5 호환되지 않는 버전으로 업데이트

- TODO

---

### 9.4.6 벤더링

- TODO

---

### 9.4.7 pkg.go.dev

- TODO

---

### 9.4.8 추가 정보

- TODO

---


