# Go_learn

Типы данных:
    Целочисленные - int,uint +(8,16,32,64),intprt,byte
    С пл.точкой - float32, float64
    Комплексные - complex64,complex128
    Булевые - bool
    Строка - string
    Константы
    Массивы:
        array - последовательность фиксированной длины
        slice - не массив, а указатель–длина–вместимость(последовательность динамической длины)
    Мапы - массив "ключ-значение"
    
Циклы
    for,forr -  
        for инициализацияСчетчика; условие; изменениеСчетчика {
            // действия
        }
        H:
            for i := 0; i < 6; i++ {
               // действия
            }
            

    foreach - for index, value := [массив] { тело цикла }
        H:
            var arr = string[4] {"a","b","c","d"}
            for indx, item := range arr {
                fmt.Printf("%d. %d\n", indx, item)
            }

    do {} while
        бесконечный (do {} while(true)) - 
            for { 
                 // действия
             }
        с условиями
            1. 
                for {
                    тело цикла
                    if(условие)
                        break;
                }
            2.
                var work = false;
                for ok := true; ok; ok = !work {
                    тело цикла
                }

функции
    func NameFunc([a type]) [type return]{
        тело функции
    }     
        