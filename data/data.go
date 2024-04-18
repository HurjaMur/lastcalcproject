package data

import (
	"sync"
)

// ID представляет структуру для хранения идентификатора и результата выражения.

// Псевдо база данных.
var Mutex = &sync.Mutex{}
var ExpressionResultMap = make(map[string]int)

// Сохраняем в базу данных
func SaveToDatabase(expression string, result int) {
	Mutex.Lock()
	defer Mutex.Unlock()
	ExpressionResultMap[expression] = result
}
