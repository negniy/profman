# Profman

Утилита CLI для управления YAML‑профилями в **текущей директории**.

Профиль — это файл `<name>.yaml` со структурой:

```yaml
user: exampleUser
project: exampleProject
```

### Использованные библиотеки

При разработке использовалась библиотека **urfave/cli v3**, потому что:

- Нужны подкоманды (`profile create/get/list/delete`) — библиотека это поддерживает «из коробки».
- API простой и декларативный: команды и флаги описываются удобно и читаемо.
- Возможностей больше, чем у стандартного пакета `flag`, но при этом она легче и проще в освоении, чем Cobra.
- Легко настраивается вывод справки и шаблонов.
- У библиотеки хорошая документация и активное сообщество, поэтому быстро находятся примеры и ответы.


## Установка

### Вариант 1: локальная сборка

```bash
go build -o profman ./cmd/profman
./profman profile list
```

### Вариант 2: `go install` (чтобы запускать из любого места)

```bash
go install ./cmd/profman
```

Бинарник попадёт в `$GOBIN` (или `$GOPATH/bin`), убедись, что эта папка в `PATH`. Тогда можно просто:

```bash
profman profile create --name=test --user=alice --project=demo
```

## Использование

```bash
# Создать профиль
profman profile create --name=test --user=alice --project=demo

# Показать профиль
profman profile get --name=test

# Список всех профилей (*.yaml в текущей папке)
profman profile list

# Удалить профиль
profman profile delete --name=test

# Помощь
profman profile -h
profman profile create -h
```
