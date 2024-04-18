package main

import (
	"fmt"
	"time"

	"github.com/HurjaMur/lastcalcproject/data"
	"github.com/HurjaMur/lastcalcproject/orchestrator"
)

func main() {
	var expression string
	var timee int
	fmt.Print("Введите ваше выражение")
	x, _ := fmt.Scan(&expression) // Введите ваше выражение
	fmt.Print("Введите время операции")
	y, _ := fmt.Scan(&timee)                         // Введите время обработки операции
	orch := orchestrator.NewOrchestrator(expression) //Добавляем наше выражение в оркестратор
	exp, _ := orch.ParseAndProcessExpression()
	time.Sleep(time.Duration(timee) * time.Second) // Введите сколько времени должно пройти до вычисления
	x += y                                         //Просто юзаем, чтобы не было ошибок в компиляторе
	res := data.ExpressionResultMap[fmt.Sprint(expression)]
	fmt.Printf("Корректность вашего выражения: %v, Ответ на ваше выражение: %v\n", exp, res)
}
