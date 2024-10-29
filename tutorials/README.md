---
title: Go Tutorials
url: https://go.dev/doc/tutorial/
---

## Endpoints

Here are the endpoints you’ll create in this tutorial.

**/albums**

- GET – Get a list of all albums, returned as JSON.
- POST – Add a new album from request data sent as JSON.

**/albums/:id**

- GET – Get an album by its ID, returning the album data as JSON.

## How to Run.

```bash
go run .
```

## How to Test.

Get All Albums.

```bash
curl http://localhost:8080/albums
```

Add Album.

```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

Get Album by ID.

```bash
curl http://localhost:8080/albums/4
```
