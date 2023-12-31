## Мікросервіс пакетів
У даному розділі описана вся потрібна інформація про мікросервіс пакетів.

***
### Завантаження пакету на сервер
**Адреса:** http://host/upload-pkg/ <br>
**Поля форми(регістр важливий):**
* _files_info(text)_ - Дане поле використовує у якості значення текст в форматі json. Текст описує структуру директорій, тобто,
назву(шлях), файли та вкладені директорії. Дані операції можна здійснити за допомогою пакету 
<u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.
Нижче наведено приклад вмісту даного поля 
```
{
  "name": "tee",
  "fileNames": ["file0.txt", "file00.txt", "file000.txt"],
  "innerDir": [
  {
    "name": "tee/dir1",
    "fileNames": ["file1.txt", "file11.txt"]
  },
  {
    "name": "tee/dir2",
    "fileNames": [],
    "innerDir": [
        {
            "name": "tee/dir2/innerD",
            "fileNames": ["file1.txt"]
        },
        {
            "name": "tee/dir2/innerD2",
            "fileNames": ["file1.txt", "file11.txt"]
        }
    ]
  }]
}
```
* _username(text)_ - Ім'я користувача
* _password(text)_ - Пароль
* _version(text)_ - Версія пакету
* _DIR$\<dirname>(files)_ - це особливе поле, тому що воно має динаму назву. Назва поля відповідає назві каталога з
файлами описаному у полі **files_info**. Кількість даних полів повинна дорівнювати кількості каталогів з файлами.
Виходячи з прикладу структури пакету який наведений вище назви полів повинні бути наступні **DIR$tee**, **DIR$tee** ...
Відповідно значення перого поля повинно містити файли "file0.txt", "file00.txt", "file000.txt", а значення другого поля
"file1.txt", "file11.txt".<br>
Якщо каталог не містить фалів його **не потрібно** надсилати в подібному форматі, інформації у **files_info** вистачить.<br>
Для зручності рекомендується використовувати пакет <u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Перевірка наявності пакету
**Адреса:** http://host/exists-package/ <br>
**Поля форми(регістр важливий):**

* _username(text)_ - Ім'я користувача
* _pkgName(text)_ - Назва пакету
* _version(text)_ - Версія пакету

**Відповідь:** <u>baseproto.BaseResponse</u>. Відповідь містить поточну інформацію про стан запиту.

***
### Завантаження пакету із сервіса
**Адреса:** http://host/download-package/ <br>
**Поля форми(регістр важливий):**

* _username(text)_ - Ім'я користувача
* _pkgName(text)_ - Назва пакету
* _version(text)_ - Версія пакету

**Відповідь:** Дана команда у відповідь відправляє форму, яка відформатована як для [завантаження пакету на сервер](#завантаження-пакету-на-сервер).
Тобто, вона містить поле **files_info** та вже знайомі поля **DIR$**. Усі дані, які потрібні для створення пакету на машині 
клієнта передаються у формі, це означає, що клієнт буде повинен зробити схожі дії, що і сервер підчас завантаження пакету в базу даних.<br>
Якщо команда операція на сервері по будь-якій причині завершится не так як очікувалось буде повернуто вже знайома відповідь
<u>baseproto.BaseResponse</u>.

***
### Показ усіх версій пакету
**Адреса:** http://host/package-versions/ <br>
**Поля форми(регістр важливий):**
* _username(text)_ - Ім'я користувача
* _pkgName(text)_ - Назва пакету

**Відповідь:** Мікросервіс повертає <u>pkgproto.MutateShowVersionBaseResponse</u>. Api сервер повертає форму із єдиним текстовим полем
**versions**. Дані операції можна здійснити за допомогою пакету <u>[github.com/uwine4850/strux_api/pkg/uplutils](https://github.com/uwine4850/strux_api/blob/master/pkg/uplutils/upload_package.go)</u>.