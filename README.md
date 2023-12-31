# Как выполнить домашнее задание
1. Скачать этот репозиторий локально с помощью `git clone`
2. Изменить remote на свой личный репозиторий с помощью `git remote set-url`
3. Перейти в ветку домешнего задания: `git switch homework/$NUM-$NAME`
4. Скачать зависимости: `go mod tidy`
5. Как правило, в ветках есть файл `main.go`, в котором нужно написать заранее объявленную функцию.
6. Как правило, в ветках есть `main_test.go`, в который можно заглядывать.
7. Написать решение, запустить тесты: `go test ./...`
8. Если тесты PASS, то создать коммит и `git push -u origin homework/$NUM-$NAME`
9. Сделать Merge Request ветки homework/$NUM-$NAME в master

# Как проверяются задания

После шага №8 списка выше, на стороне репозитория GitLab запускаются автоматическое тестирование выполненного задания. Результат тестирования появляется в виде индикатора Pipeline на главной странице ветки, напротив Commit ID, или в разделе меню репозитория CI/CD -> Pipelines.

Если автотесты выполнились успешно, то задание попадает в Code Review, после чего ему ставится оценка.

## Условия оценок

1. Прошло CI/CD, Code Review, выполнено во время Deadline: 1.
2. Прошло CI/CD, Code Review, выполнено после Deadline: 0.75.
3. Не прошло CI/CD: 0.
4. Прошло CI/CD, не прошло Code Review: обратная связь от ревьювера в комментариях к Merge Request.

Максимальный балл по домашним заданиям: 12

Минимальный: 9