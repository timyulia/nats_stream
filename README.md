# nats_streaming
Чтобы запустить:
```
docker-compose up
cd publisher && go run main.go
cd sub && go run cmd/main.go
```
В publisher необходимо указывать полный путь к файлу или model.json, лежащий в данной директории.
Интерфейс доступен по адресу: http://localhost:8000/order/