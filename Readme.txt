Программа для просмотра данных из 1С в браузере.
Сделал Никитин Александр
Skype: Travianbot

Программа предназначена для просмотра данных из 1С в браузере,
база 1С должна быть опубликована на веб-сервере.

Порядок работы:
1. Открыть в 1С обработку:
РедактированиеСоставаСтандартногоИнтерфейсаOData.epf 
Добавить в список нужные объекты 1С: Документы, справочники и др.
которые будет видно через интерфейс OData
2. Изменить файл с настройками: settings.ini 
В котором настроить параметры доступа в базу 1С через веб.
URL=http://192.168.1.1/Baza1
Login=НикитинА
Password=
Port=8080
3. Запустить программу: 1C_Odata.exe
4. Открыть в браузере URL: http://localhost:8080
откроется страница как на скрине,
можно кликнуть на нужный вид справочника 
и все данные отобразятся в браузере в формате JSON.
Потом доделаю отображение чтоб было красивее в виде таблицы.

Сделал в рамках изучения языка Golang.
Код открыт, выложу.

Язык русский:
Тестировал: 1С:Предприятие 8.3 (8.3.16.1063)
Лицензия: Указывать имя автора и сайт.
