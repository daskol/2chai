# 2chAI

*сбор диалогов в тематических досках*

## Описание

Проект состоит из двух частей. Первая часть занимает сбором сырых диалогов из
тематических досок с тем, чтобы натренировать диалоговую модель на них. Вторая
же как раз и занимается обучением.

## Эксплуатация

На данный момент реализован консольный инструмент, позволяющий синхронизировать
базу данных нитей и постов в локальное хранилище. В качестве такого хранилища
была выбрана СУБД PostgreSQL. Модель данных не отличается от той, в которой
потребляется контент: `boards`, `threads`, `posts`. Поэтому первым этапом
развёртывания необходимо накатить миграцию выполнив что-то похожее следующую
команду.

```bash
    psql < sql/schema.sql
```

Далее, необходимо собрать инструмент синхронизации, который предоставляет
простенькое CLI для работы с API. Предполагается, что все зависимости уже
рарзрешены локально.


```bash
    go build
```

Так как CLI существенно опирается на базу данных и использует в своей работе
таблицу `boards` для того, чтобы определить следует ли следить за данной доской
или нет, следует синхронизировать список досок.

```bash
    ./2chai sync-boards
```

По-умолчанию флаг `watch` в таблице `board` не выставлен для всех досок, что
означает, что никакие данные не будут переложены в базу. Для того, чтобы
выставить флаг слежения за всеми досками, можно выполнить, например, следующее.

```bash
    psql <<< "UPDATE boards SET watch = 't';"
```

Теперь всё готово для миграции данных в локальное хранилище.

```bash
    ./2chai sync-all
```

После выполнения предыдущей команды каждый двадцать минут будет опрашиваться
API о существующих нитях и постах, которые будут сохранены в базе данных.
