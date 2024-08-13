package output

import "github.com/fatih/color"

func PrintError(value any) {
	//intVal, ok := value.(int)
	//if ok {
	//	color.Red("Код ошибки: %d", intVal)
	//	return
	//}
	//stringVal, ok := value.(string)
	//if ok {
	//	color.Red(stringVal)
	//	return
	//}
	//color.Red("Неизвестный тип ошибки: %v", value)
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки: %v", t)
	}
}

type List[T any] struct {
	elements []T
}
