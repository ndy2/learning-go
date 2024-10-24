---
title: NOTE
---

> [!NOTE]
> 서로 다른 크기로 선언된 경우 두 개의 정수는 더할 수 없다.

```go
package main

import "fmt"

func main() {
    var a int16 = 10
    var b int32 = 20

    // 명시적으로 타입을 맞춰야 한다
    fmt.Println(a + int16(b))  // int32를 int16으로 변환
}
```

---

> [!NOTE]
> 리터럴은 타입이 지정되어 있지 않다.
> 개발자가 특정 타입으로 지정하기 전까지 타입을 강제하지 않는다.

---

> [!NOTE]
> 타입 변환은 Go 가 단순성과 명확성을 대가로 약간의 자세한 정보를 추가해야 하는 선택을 한 부분이다.

