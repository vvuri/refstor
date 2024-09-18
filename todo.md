# Go миккросервис
Хранилище референсов - Reference storage

## Идея
- кидаем ссылку на картинку, добавляем текст и автора + тег - сохраняем в БД
- картинок может быть несколько
- далее может быть и видео и тайминг по времени - от и до минуты
- нейронка должна выделять базовые и вспомогательные цвета
- так же можно вытащить лут по фото
- с видео тоже можно разбить на смены планов и тайминги или на описанию происходящего

## Этапы реализации
1. **Для начала Go + Vue + SQLite**
2. Показ референсов, галерея, отображение по тегам и времени
3. Авторизация JWT
4. Микросервис - Выгрузка в PDF для дальнейшей печати
5. Микросервис анализато изображения и разбивка на компоненты 
6. Микросервис для скачивания и хранения видео
7. Docker compose или кубер
8. Профилирование и оптимизация
9. Сторедж хранилище под видео данные 
10. Деплой на прод
11. Backup данных

## За основу берем 
- связка go + ui: https://www.youtube.com/watch?v=lNd7XlXwlho
- используем vue:
- UI можно попробовать смоделировать в нейронке https://shots.so/

## Steps
1. Создание репозитория
2. Перенос шаблона прокта
3. Go + Chi - https://github.com/go-chi/chi
4. $ go mod init refstor
5. $ go mod vendor
6. $ go get github.com/lpernett/godotenv
7. $ go get -u github.com/go-chi/chi/v5
8. $ go get github.com/go-chi/render
9. Create Vue Client
- $ npm create vue@latest
- $ cd client
- $ npm install
- $ npm run dev
- Open http://localhost:5173/
- Test API in windows: ```Invoke-WebRequest -Uri http://localhost:5001/api/img/234 -Method "PUT"```
10. Redis install:
- $ choco install redis-64 --version 3.0.503 - не работает
- WSL + $ sudo apt-get install redis
- получаем  /usr/bin/redis-server 127.0.0.1:6379
- $ redis-cli  +  ping
- $ go get github.com/redis/go-redis/v9
11. UUID
- $ go get github.com/google/uuid
12. Test Created request:
- $ curl -X POST -d '{}' 192.168.1.110:5001/api/img
  {"uuid":"a15bd769-649d-451c-b707-17631744aae6","description":"","small_img":null,"date":"2024-09-16T21:19:12.5215519Z","url":""}
- $ redis-cli KEYS '*'
  1) "images"
  2) "image:323842323131354541353232"
- $ redis-cli
  127.0.0.1:6379> GET "image:323842323131354541353232"
  "{\"uuid\":\"a15bd769-649d-451c-b707-17631744aae6\",\"description\":\"\",\"small_img\":null,\"date\":\"2024-09-16T21:19:12.5215519Z\",\"url\":\"\"}"
13. Test create with param
- $ curl -X POST -d '{"description":"text for link","link":"ya.ru/img1"}' 192.168.1.110:5001/api/img
  {"uuid":"36d4893f-95c4-4f66-819f-fa943365a496","description":"text for link","small_img":null,"date":"2024-09-16T22:26:12.1160674Z","url":"ya.ru/img1"}
- $ redis-cli
  127.0.0.1:6379> keys *
  127.0.0.1:6379> GET "image:304637463630413046343745"
  "{\"uuid\":\"36d4893f-95c4-4f66-819f-fa943365a496\",\"description\":\"text for link\",\"small_img\":null,\"date\":\"2024-09-16T22:26:12.1160674Z\",\"url\":\"ya.ru/img1\"}"
14. All values in key:images
    127.0.0.1:6379> smembers "images"
    1) "image:323842323131354541353232"
    2) "image:304637463630413046343745"
15. Test List Get tequest
    $ sudo apt  install jq
    $ curl -sS 192.168.1.110:5001/api/img | jq
16. -
 
    

## Vue
1. $ npm install -g @vue/cli
2. $ npm vue --version
  10.8.3
3. $ vue --version
   @vue/cli 5.0.8
4. Add Vuetify
- $ vue add vuetify
- Error
5. Run:
- $ vue ui
- http://localhost:8000/project/select
6. 

## Redis CLI
- List all keys using the KEYS command:
  $ redis-cli KEYS '*'
- Get all keys from redis-cli
  redis 127.0.0.1:6379> keys *
- Get list of patterns
  redis 127.0.0.1:6379> keys d??
- This will produce keys which start by 'd' with three characters.
  redis 127.0.0.1:6379> keys *t*
- This wil get keys with matches 't' character in key
- Count keys from command line by
  redis-cli keys * |wc -l 
- Or you can use dbsize
  redis-cli dbsize
- Here are the commands to retrieve key value(s):
  if value is of type string -> GET <key>
  if value is of type hash -> HGET or HMGET or HGETALL <key>
  if value is of type lists -> lrange <key> <start> <end>
  if value is of type sets -> smembers <key>
  if value is of type sorted sets -> ZRANGEBYSCORE <key> <min> <max>
  if value is of type stream -> xread count <count> streams <key> <ID>. https://redis.io/commands/xread
- Use the TYPE command to check the type of value a key is mapping to:
  type <key>



