Для запуска программы требуется: 

1) Создать БД в psql и ввести название в .env (для удобства понимания осталвены локальные данные), про .env в пункте '2';
2) Ввести данные для БД в файл .env (Примичание: по умолчанию используется порт 5432)
3) Для установки зависимостей выполнить команду go mod tidy
4) Переименовать импорт локальных пакетов с tz_todo_list_1 на используемый
