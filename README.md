
# Итоговая задача 2 модуля курса по Go от Яндекса

Этот проект реализует веб-сервис, принимающий выражение через http запрос и возвращающий результат вычислений. Да, это калькулятор.


## Инструкция по запуску:

1) Скопируйте репозиторий и перейдите в место расположения проекта:

```bash
git clone https://github.com/Vyber777/CalcModule2

cd CalcModule2
```

2) Запустите оркестратор:

```bash
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

go run cmd/orchestrator.start/main.go
```

Выведется Starting Orchestrator on port 8080.

Переходим в репозиторию с проектом:

```bash
cd Module2calc
```

3) Затем запускаем agent:

```bash
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=http://localhost:8080

go run cmd/agent.start/main.go
```

Вы получите ответ:
Starting Agent...  
Starting worker 0  
Starting worker 1  
Starting worker 2  
Starting worker 3  

(Или иное количество в зависимости от числа в COMPUTING_POWER)

## Примеры использования:

Успешный запрос:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2*2+2"}'
```

Ответ:

```bash
{
  "id": "..."
}
```

После можно посмотреть этап выполнения данного запроса и его результат:

```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```

Вывод:

```bash
{"expressions":[{"id":"1740240110508066400","expression":"2*2+2,"status":"pending"}]}
```

Если вычисления выполнены то:

```bash
{"expression":{"id":"1","expression":"2*2+2","status":"completed","result":6}}
```

Или узнать точный результат нужного выражения по его id:

```bash
curl --location 'http://localhost:8080/api/v1/expressions/id'
```

## Ошибки при запросах:

### Ошибка 404 (отсутствие выражения):

```bash
{"error":"Expression not found"}
```

### Ошибка 422 (невалидное выражение):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2+a"
}'

```
Ответ:

```bash
{
  {"error":"expected number at position 2"}
}
```

### Ошибка 500 (внутренняя ошибка сервера):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2/0"
}'
```
Ответ:

```bash
{
  Worker n: error computing task 3: division by zero
}
```

## Запуск тестов агента:

1) Переходим в папку с модулем.

```bash
cd Module2calc
```

2) Запускаем тестирование:

```bash
go test ./internal/agent/agent_calculation_test.go
```

3) При успешном прохождении теста должен вывестись ответ:

```bash
ok  	calc_service/internal/evaluator	0.001s
```

4) При ошибке в тестах будет указано где она произошла и информация по ней.