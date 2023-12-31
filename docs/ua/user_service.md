## Мікросервіс користувачів
У даному розділі описана вся потрібна інформація про мікросервіс користувачів.

***
### Створення користувача
**Адреса:** http://host/create-user/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача
* _password(text)_ - Пароль

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Перевірка наявності користувача
**Адреса:** http://host/user-exist/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Видалення користувача
**Адреса:** http://host/user-delete/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача
* _password(text)_ - Пароль

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Оновлення паролю
**Адреса:** http://host/user-password-update/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача
* _password(text)_ - Пароль
* _newPassword(text)_ - Новий пароль

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Вхід користувача
**Адреса:** http://host/user-log-in/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача
* _password(text)_ - Пароль

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.
