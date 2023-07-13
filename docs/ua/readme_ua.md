<h1 align="center" style="margin: 0">Strux api</h1>

## Інші розділи
* [Структура мікросервісів](https://github.com/uwine4850/strux_api/blob/master/docs/ua/microservices_structure.md)
* [Детальніше про мікросервіс користувачів](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* [Детальніше про мікросервіс пакетів](https://github.com/uwine4850/strux_api/blob/master/docs/ua/package_service.md)

## Про проєкт
Strux api це проєкт призначений для керування пакетами(також для роботи потрібен
[клієнт серверу](https://github.com/uwine4850/strux)). Пакет - це
лише директорія із проєктом користувача. Все, що знаходиться в одній
директорії при належній підготовці можна вважати пакетом. Щоб інші користувачі
могли завантажувати пакети їх потрібно перемістити на сервер за допомогою
клієнту. Для створення пакету потрібно використати [клієнт серверу](https://github.com/uwine4850/strux).<p></p>

## Початок роботи
Для початку роботи потрібно виконати описану нижче інструкцію.

Якщо даний проєкт використовується у якості серверу репозиторій потрібно завантажити за даним посиланням:<br>
```
https://github.com/uwine4850/strux_api
```
Для запуску сервера потрібно запустити сервер. Для початку потрібно запустити мікросервіси(main.go) у каталозі services:<br>
* Мікросервіс користувачів
* Мікросервіс пакетів
<p></p>
Після цього потрібно запустити http сервер який знаходиться з аадресою cmd/main.go.

Якщо даний проєкт використовується у якості додаткового пакету потрібно його встановити за допомогою команди<br>
```
go get github.com/uwine4850/strux_api
```

## Інформація про api
Сервіс містить наступні команди:
* http://host/create-user/ - створення користувача. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* http://host/user-exist/ - перевірка наявності користувача. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* http://host/user-delete/ - видалення користувача. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* http://host/user-password-update/ - оновлення паролю. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* http://host/user-log-in/ - вхід користувача у систему. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/user_service.md)
* http://host/upload-pkg/ - завантаження пакету на сервер. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/package_service.md)
* http://host/exists-package/ -  перевірка наявності пакету. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/package_service.md)
* http://host/download-package/ - завантаження пакету із сервера. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/package_service.md)
* http://host/package-versions/ - повертає усі доступні версії пакету. [[Info]](https://github.com/uwine4850/strux_api/blob/master/docs/ua/package_service.md)

### Коротко про алгоритм роботи
Щоб виконати команду потрібно відправити на одну з адрес http форму. Наприклад, щоб відправити форму за допомогою golang
можна використати http.NewRequest. Сервер приймає лише дані у вигляді форми, поля можуть містити текстові або файлові дані.
Після опрацюванння форми сервер повертає або текстові дані(text/plain) або форму (multipart/form-data). Для використання
api краще викоритовувати клієнт але це не обов'язково.