package controller

func createTodo() {
	v, err := model.createTodo()
	if err != nil {

	}
}

func updateTodo() {
	v, err := model.updateTodo()
	if err != nil {

	}
}

func deleteTodo() {
	err := model.deleteTodo()
	if err != nil {

	}
}

func getTodos() {
	v, err := model.getTodos()
	if err != nil {

	}
}
