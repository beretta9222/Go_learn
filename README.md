# Go_learn

Типы данных:
    Целочисленные - int,uint +(8,16,32,64),intprt,byte
    С пл.точкой - float32, float64
    Комплексные - complex64,complex128
    Булевые - bool
    Строка - string
    rune - это инструмент для работы с символами Unicode. rune явно говорит, что мы работаем не просто с числом, а с кодовой точкой Unicode
    Константы
    Массивы:
        array - последовательность фиксированной длины
        slice - не массив, а указатель–длина–вместимость(последовательность динамической длины)
    Мапы - массив "ключ-значение"
    Кастомный тип:
        H:  
            type age int
            type sex string
            type money n

Cтруктура:
    набор различных типов данных struct{}{}
        H: type people struct {
                age int
                sex string
            }{}
    у структуры могут быть методы 

    func ([имя перемнной] [название структуры]) НазваниеМетода(параметры метода) возвращаемый тип {
        //телометода
    }
        H: func (p people) getInfo() string {
            return fmp.Sptrint("age: %s\n sex: %s");
        }



Интерфейсы:
    Интерфейсы представляют абстракцию поведения других типов. 
    Интерфейсы позволяют определять функции, которые не привязаны к конкретной реализации. То есть интерфейсы определяют некоторый функционал, но не реализуют его.

    type имя_интерфейса interface{
        определения_функций
    }
    Н: type vehicle interface{
            move()
        }

        type Vehicle interface{
            move()
        }
        
        func main() {
            
            var tesla Vehicle   // переменная интерфейса
            fmt.Println(tesla)  // nil
        }

    Реализация интерфейса

        type Vehicle interface{
            move()
        }
        
        // структура "Автомобиль"
        type Car struct{ }
        
        // структура "Самолет"
        type Aircraft struct{}
        
        
        func (c Car) move(){
            fmt.Println("Автомобиль едет")
        }
        func (a Aircraft) move(){
            fmt.Println("Самолет летит")
        }
        
        func main() {
            
            var tesla Vehicle = Car{}
            var boing Vehicle = Aircraft{}
            tesla.move()    // Автомобиль едет
            boing.move()    // Самолет летит
        }

    Полиморфизм
        Интерфейсы в языке Go позволяют реализовать концепцию полиморфизма - способность принимать многообразные формы:
        есть несколько типов, которые имеют одни и те же методы интерфейса, но по разному их реализуют. 
        То есть имеется один и тот же функционал и множество форм его реализации, и поведение интерфейса изменяется в соответствии с типом, который реализует интерфейс


        type Vehicle interface{
            move()
        }
        
        type Car struct{ model string}
        type Aircraft struct{ model string}
        
        
        func (c Car) move(){
            fmt.Println(c.model, "едет")
        }
        func (a Aircraft) move(){
            fmt.Println(a.model, "летит")
        }
        
        func main() {
            
            tesla := Car{"Tesla"}
            volvo := Car{"Volvo"}
            boeing := Aircraft{"Boeing"}
            
            vehicles := [...]Vehicle{tesla, volvo, boeing}
            for _, vehicle := range vehicles{
                vehicle.move()
            }
        }

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
    switch value := значение; value {
        case "значение1":
            // действия
        case "значение2":
            // действия
        default:
            // действия
    }

функции
    func NameFunc([a type]) [type return]{
        тело функции
    }     
        