# Otus Microservice architecture Homework 2

## Домашнее задание выполнено для курса ["Microservice architecture"](https://otus.ru/lessons/microservice-architecture/)

Для запуска использовать команду

```bash
kubectl apply -f deployments
```

Тесты Postman расположены в директории `test/postman`. Запуск тестов.

```bash
newman run ./test/postman/test.postman_collection.json
```

Или с использованием Docker.

```bash
docker run -v $PWD/test/postman/:/etc/newman --network host -t postman/newman:alpine run test.postman_collection.json
```
