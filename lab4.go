package main

// подключаем пакет для вывода информации
import "fmt"

// Задаём структуру Token
type Token struct {
    // Сообщение
    data string
    // Получатель
    recipient int
    // Время жизни сообщения
    ttl int
}

// Глобальная переменная - количество узлов
var N int

func main() {
    // Создаём сообщение в виде экземпляра структуры Token
    t := Token{data: "message", recipient: 2, ttl: 14}
    // Создаём channel - Канал для связи узлов
    c := make(chan string)
    // Начинаем отправку с узла №i
    i := 7
    // Запускаем N потоков (goroutine)
    N = 10
    go thread(i, c, t)
    // Когда в канале появится информация - считываем её
    fmt.Println(<-c)
}

func thread(i int, c chan string, t Token) {
    // Если текущий узел является получателем токена
    if (i == t.recipient) {
        c <- "Token is received!"
    // Иначе если время жизни сообщения закончилось
    } else if (t.ttl == 0) {
        c <- "Time is over!"
        // Иначе передаём следующему узлу
        } else {
            // При этом уменьшаем время жизни сообщения
            t.ttl -= 1
            // Задаём следующий узел (если прошли круг, то снова рассматриваем узел №1
            i++
            if (i > N) {
                i = 1
            }
            // Запускаем следующий поток
            go thread(i, c, t)
        }
}