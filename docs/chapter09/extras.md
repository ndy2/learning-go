---
title: 번외
aliases:
  - Extras
---

## Go Modules Redirection (?) <- 진짜 있는건지 확인 필요

Q: 모듈의 이름이 github.com 으로 시작하는데 github.com 이 아닌 gitlab 으로 관리할 수 있니?

A: 

네, Go 모듈의 이름이 github.com으로 시작하더라도 실제로는 GitLab과 같은 다른 Git 저장소에서 관리할 수 있습니다. 다만, 모듈 경로와 저장소 URL이 일치해야 하므로, 몇 가지 설정이 필요합니다.

Go는 기본적으로 모듈 이름과 일치하는 URL에서 코드를 다운로드하려고 시도합니다. 예를 들어, 모듈 이름이 github.com/username/repo라면, Go는 github.com/username/repo에서 해당 코드를 다운로드하려고 합니다. 하지만 gitlab.com과 같은 다른 저장소를 사용할 수 있도록, **Go Modules 는 리디렉션**을 허용합니다.


