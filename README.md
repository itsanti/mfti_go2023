Ваша задача написать локальный REST API сервер с локальной базой данных PostgreSQL для хранения покемонов. У сервера должен быть всего один эндпоинт `/pokemons`. Этот эндпоинт должен поддерживать следующие операции:
- Добавить покемона
- Получить список всех покемонов
- Получить покемона по его ID

# Требования

- https://github.com/julienschmidt/httprouter
- https://github.com/go-gorm/gorm
- Структура Pokemon из [homework/4](https://gitlab.atp-fivt.org/courses-public/golang/homeworks/-/tree/homework/4-pokemon-api-client)
- Возвращать правильные HTTP статусы (см. тесты)
- Сервер слушает порт 8080

# Рекомендации

- Лучше всего использовать [Docker Compose](https://docs.docker.com/compose/) для поднятия локальной БД.
- В `.gitlab-ci.yml` есть команды для поднятия локальной БД (Linux)