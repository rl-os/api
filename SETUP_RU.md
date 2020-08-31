# Установка и настройка сервера

## Требования

Для работы сервера необходимо

* GNU Make >= 4.2.1
* Golang >= 1.13
* NodeJS >= 10 с yarn или npm 3.
* Python 3.7+ (т.к. используется в коде async/await)
* Redis
* RabbitMQ
* PostgreSQL 11+
* Reverse proxy (nginx/traefik/haproxy/что то еще) с самоподписанным сертификатом

Или

* Brain
* Docker Engine => 17.12.0
* Docker Compose

## Локальная установка для разработки

Для того что бы скомпилировать сервер потребуется GoLang версии 1.13 и выше.
```bash
# через GVM (https://github.com/moovweb/gvm)
$ bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
$ gvm install -B go1.15 # или выше
$ gvm use go1.15 --default

# через brew
$ brew install golang
```

## Production Ready установка

TODO
