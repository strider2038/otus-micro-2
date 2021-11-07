# Otus Microservice architecture Homework 2

## Домашнее задание выполнено для курса ["Microservice architecture"](https://otus.ru/lessons/microservice-architecture/)

Для запуска использовать команды

```bash
## для первого запуска зарегистрировать репозиторий bitnami
helm repo add bitnami https://charts.bitnami.com/bitnami
## запуск БД
helm install --wait -f deployments/postgres/values.yaml postgres bitnami/postgresql
## запуск проекта
helm install --wait -f deployments/user-service-values.yaml user-service ./deployments/user-service
```

Тесты Postman расположены в директории `test/postman`. Запуск тестов.

```bash
newman run ./test/postman/test.postman_collection.json
```

Или с использованием Docker.

```bash
docker run -v $PWD/test/postman/:/etc/newman --network host -t postman/newman:alpine run test.postman_collection.json
```
