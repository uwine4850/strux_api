## Структура мікросервісів
* Усі мікросервіси зберігаються у каталозі services/
* Файли protobuf знаходяться у каталозі services/proto_source
* У каталозі services/protofiles зберігаються згенеровані файли
* Файл utils відповідає за загальну логіку сервісів
* Кожен мікросервіс містить файл main.go. У ньому ініціалізується відповідний сервіс, а при запуску файлу запускається і
сервіс
* У каталозі services/<service_name>/internal зберігаються файли з обробниками команд. У кажному файлі зберігається логіка
однієї команди. Відповідна команда викликається у файлі main.go у відповідному методі
* В загальному обробка даних у сервісах проходить наступним чином підключення до бази даних -> отримання структури db.DatabaseOperation{}
для взаємодії із вибраною таблицею -> виконання бізнес логіки -> повернення відповіді.