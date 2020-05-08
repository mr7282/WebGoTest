package main

import(
	"homework-3/server"
	"os"
	"os/signal"
)


// 1. Создайте роут и шаблон для отображения всех постов в блоге.
// 2. Создайте роут и шаблон для просмотра конкретного поста в блоге.
// 3. Создайте роут и шаблон для редактирования и создания материала.
// 4. * Добавьте к роуту редактирования и создания материала работу с Markdown с помощью пакета blackfriday.
// Рекомендуем хранить контент поста в блоге в типе template.HTML, чтобы использовать html-разметку внутри поста (для blackfriday это обязательное условие корректного отображения материала).




func main() {
	go func ()  {
		server.StartServer()
	}()
	stopSig := make(chan os.Signal)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig
}


