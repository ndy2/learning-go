---
title: "Tutorial: Accessing a relational database"
url: https://go.dev/doc/tutorial/database-access
---

## setup mysql

```
k create -f mysql.pod.yaml
k exec -it mysql-pod -- mysql -u root -p
```
