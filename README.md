# Проект 1 модуля яндекс лицея

## Запуск

```cmd
go run ./cmd/api -p=8080
```
флаг `p` отвечает за используемый порт

# Примеры запросов
## 200 запрос
```bash
curl --location '127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}' \
-X POST
```

## 422 запрос
```bash
curl --location '127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{"expression": "2+2*2zxc"}' \
-X POST
```