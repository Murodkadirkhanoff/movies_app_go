# Проект на Go с Docker и PostgreSQL

Этот проект представляет собой API для работы с фильмами и пользователями, использующие PostgreSQL в качестве базы данных. Проект настроен с использованием Docker для удобства развертывания.

## Содержание
- [Docker конфигурация](#docker-конфигурация)
- [API Endpoints](#api-endpoints)

## Docker конфигурация

Для запуска проекта с помощью Docker, вам нужно выполнить несколько шагов:

1. **Склонируйте репозиторий:**
    ```bash
    git clone https://github.com/Murodkadirkhanoff/movies_app_go.git
    cd movies_app_go
    ```

2. **Создайте файл `.env`:**
    В корне проекта создайте файл `.env` и добавьте следующие переменные:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    ```

3. **Запустите проект с помощью Docker Compose:**
    В корне проекта выполните команду:
    ```bash
    docker-compose up -d --build
    ```

    Это создаст два контейнера: один для приложения (Go) и второй для базы данных (PostgreSQL).

4. **Доступ к PgAdmin:**
    PgAdmin будет доступен по адресу `http://localhost:5050`. Логин и пароль по умолчанию:
    - Email: `admin@example.com`
    - Password: `admin123`

5. **Доступ к API:**
    После успешного запуска приложения, API будет доступно по адресу `http://localhost:8080`.

## API Endpoints

API предоставляет несколько конечных точек для работы с фильмами:

### 1. Создание нового фильма
- **URL:** `/movies`
- **Метод:** `POST`
- **Тело запроса:**
    ```json
    {
        "title": "Название фильма",
        "director": "Режиссер",
        "year": 2023,
        "plot_id": 1
    }
    ```
- **Ответ:**
    - **201 Created:** Если фильм успешно создан.
    - **400 Bad Request:** Если данные запроса неверны.

### 2. Получение списка всех фильмов
- **URL:** `/movies`
- **Метод:** `GET`
- **Ответ:**
    ```json
    [
        {
            "id": 1,
            "title": "Название фильма",
            "director": "Режиссер",
            "year": 2023,
            "plot_id": 1,
            "plot": {
                "id": 1,
                "title": "Описание"
            }
        }
    ]
    ```

### 3. Получение фильма по ID
- **URL:** `/movies/:id`
- **Метод:** `GET`
- **Ответ:**
    ```json
    {
        "id": 1,
        "title": "Название фильма",
        "director": "Режиссер",
        "year": 2023,
        "plot_id": 1,
        "plot": {
            "id": 1,
            "title": "Описание"
        }
    }
    ```
    - **404 Not Found:** Если фильм с таким ID не найден.

### 4. Обновление информации о фильме
- **URL:** `/movies/:id`
- **Метод:** `PUT`
- **Тело запроса:**
    ```json
    {
        "title": "Новое название фильма",
        "director": "Новый режиссер",
        "year": 2024,
        "plot_id": 2
    }
    ```
- **Ответ:**
    - **200 OK:** Если фильм успешно обновлен.
    - **404 Not Found:** Если фильм с таким ID не найден.

### 5. Удаление фильма
- **URL:** `/movies/:id`
- **Метод:** `DELETE`
- **Ответ:**
    - **200 OK:** Если фильм успешно удален.
    - **404 Not Found:** Если фильм с таким ID не найден.

---