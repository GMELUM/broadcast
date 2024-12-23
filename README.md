# Broadcast Program Documentation

## 1. Компиляция программы

В зависимости от платформы, на которой вы хотите запустить программу, используйте одну из следующих команд для компиляции:

- **Linux**:
  ```bash
  CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/broadcast .
  ```

- **macOS**:
  ```bash
  CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/broadcast .
  ```

- **Windows**:
  ```bash
  CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/broadcast.exe .
  ```

Компиляция создаст исполняемый файл в директории `dist`.

---

## 2. Подготовка `.csv` файла

Рядом с исполняемым файлом необходимо разместить `.csv` файл с данными в следующем формате:

```csv
id
1093776793
```

Пример файла можно назвать `data.csv`.

---

## 3. Описание сообщения

Для работы программы требуется создать JSON-файл с описанием сообщения. Вот пример формата:

```json
{
    "image": "https://api.ratingtma.com/cdn/post.jpeg",
    "text": "🎄✨ Holiday Magic is Here, Shining Stars ✨🎄\n\n✍️ Collect toys from tasks, mini-games, and the shop.\n🧥  Decorate your tree to unlock rewards.\n🤑 Earn passive income with every toy!\n\nStart celebrating and earning now! 🚀",
    "effect_id": "5104841245755180586",
    "keyboard": [
        {
            "Play": "http://t.me/rating/app"
        }
    ]
}
```

### Поля JSON:
- **image**: URL изображения для отправки.
- **text**: Текст сообщения.
- **effect_id**: Идентификатор эффекта.
- **keyboard**: Массив кнопок с названиями и ссылками.

---

## 4. Запуск программы

После подготовки всех файлов можно запустить программу:

1. **Linux/macOS**:
   ```bash
   ./dist/broadcast
   ```

2. **Windows**:
   ```bash
   dist\broadcast.exe
   ```

При запуске программа запросит:

1. **Токен бота** — укажите токен вашего Telegram-бота.
2. **RPS (Requests Per Second)** — укажите количество запросов в секунду для рассылки.

---

## Пример запуска:
- **Linux**:
  ```bash
  ./dist/broadcast
  ```

**Ввод параметров:**
- Введите токен бота: `123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11`
- Введите RPS: `5`

---

## 5. Дополнительные советы

- Убедитесь, что ваш `.csv` файл корректно структурирован и находится в одной папке с исполняемым файлом.
- Используйте JSON-валидатор для проверки вашего файла с сообщением.

---

Теперь вы готовы к использованию программы! 🎉
