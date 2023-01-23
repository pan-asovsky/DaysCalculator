### DaysCalculator
#### 1. Описание
Web-приложение на фреймворке gin, c обработчиком маршрута `/when/:year`, который указывает, сколько дней осталось или прошло c 1-го января года, указанного в параметре маршрута, и middleware проверяющий наличие заголовка `X-PING`
- Например: `/when/2000` покажет сколько дней прошло с 01.01.2000
- Например: `/when/2030` покажет сколько дней осталось до 01.01.2030

*Для избежания путаницы указываю, что приложение не будет учитывать сегодняшний день в расчёте, считая его неполным. Так для 20.01.2023 при текущем 2023 году, количество прошедших дней будет указано как 19. Сделано из субъективных соображений о том, как правильно считать дни.*

#### 2. Запуск
- Компиляция исполняемого файла `make build`
- Запуск исполняемого файла `./main`
- Запуск без компиляции файла `make run`
- Запуск тестов `make test`