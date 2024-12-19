
# web-calculator

Web-calculator представляет из себя веб-сервис, позволяющий пользователю отправлять cURL запросы с арифметическими выражениями и получать в ответ его результат.



## Функционал

- Поддерживаются операции сложения, вычитания, умножения и деления
- Выражение может вводиться как с пробелами, так и без
- Поддерживаются выражения, состоящие из положительных натуральных чисел


## Требования

- Go версии ```1.20``` или новее


## Установка

1. Клонирование репозитория

```bash
git clone https://github.com/bulbosaur/web-calculator-golang
```

2. Запуск сервера
   
```bash
cd web-calculator-golang
go run cmd/main.go
```


## Переменные окружения

| Переменная | Описание | Значение по умолчанию |
|------------|----------|-----------------------|
| ```PORT``` | Порт для запуска сервера | ```8080``` |
| ```HOST``` | Хост для запуска сервера | ```localhost``` |



## API

Базовый URL по умолчанию: ```http://localhost:8080```

| API endpoint | Ответ сервера | Код ответа |
|--------------|---------------|------------|
| ```/api/v1/calculate``` | json ```{"result":"результат выражения"}``` | 200 |
| ```/coffee``` | ```I'm a teapot``` | 418 |

### Коды ответов
- 200 - Успешный запрос
- 400 - Некорректный запрос
- 403 - Доступ запрещен
- 404 - Ресурс не найден
- 422 - Некорректное выражение (например, буква английского алфавита вместо цифры)
- 500 - Внутренняя ошибка сервера