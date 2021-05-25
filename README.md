# Приложение "Список дел" (только REST API часть)

Запуск контейнера с postgresql (для простоты MVP база в нем):
```
docker pull postres
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
docker exec -it 5a9d643bdaf3 /bin/bash
```
Создадим файлы миграций для БД:
```
brew install golang-migrate
migrate create -ext sql -dir ./schema -seq init
```
В них опишем создание-удаление таблиц (директория ./schema в проекте).  
Запускаем миграции:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
```
В контейнере можно убедиться, что БД с таблицами созданы:
```
docker ps
docker exec -it PID /bin/bash

psql -U postgres
\d
select * from users;
```

Запуск приложения:
```
go run ./cmd/main.go
```
  
  
Документация методов:  
http://localhost:8000/swagger/index.html
