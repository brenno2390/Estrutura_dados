/*Exemplo 01*/
package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a

	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de p (endereço de a):", p)
	fmt.Println("Valor apontado por p:", *p)
	modify(10)
	sub(100)
	modify1(150)
	passagem(200.54)
	passagem2(200.54)
	swap(10, 20)

	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("Lista Encadeada Simples:")
	list.Print()
	list.RemoveFirst()
	fmt.Println("Após remover o primeiro:")
	list.Print()
}

/* Exemplo 02*/

func modify(val int) {
	var a int = 42
	var p *int = &a
	fmt.Println("Valor de a antes da modificação:", a)
	*p = 100
	fmt.Println("Valor de a após a modificação:", a)

}

/*Exercicio 01*/

func sub(val int) {
	var b float64 = 42.13
	var p2 *float64 = &b
	fmt.Println("Valor de b:", b)
	fmt.Println("Endereço de b:", &p2)
}

func modify1(val int) {
	var b float64 = 42.13
	var p2 *float64 = &b
	fmt.Println("Valor de a antes da modificação:", b)
	*p2 = 150.13
	fmt.Println("Valor de a após a modificação:", b)

}

/*Exemplo 03*/

func increment(val int) {
	val++
	fmt.Println("Dentro da função increment:", val)
}

func passagem(val float64) {
	x := 10
	increment(x)
	fmt.Println("Fora da função increment:", x)
}

/*Exemplo 04*/
func incrementByPointer(val *int) {
	*val++
	fmt.Println("Dentro da função incrementByPointer:", *val)
}
func passagem2(val float64) {
	x := 10
	incrementByPointer(&x)
	fmt.Println("Fora da função incrementByPointer:", x)
}

/*Exercicio 02*/

func swap(a, b int) {
	var p1 *int = &a
	var p2 *int = &b
	fmt.Println("Valor de a antes da troca:", a)
	fmt.Println("Valor de b antes da troca:", b)
	aux := *p1
	*p1 = *p2
	*p2 = aux
	fmt.Println("Valor de a após a troca:", a)
	fmt.Println("Valor de b após a troca:", b)
}

/* Lista Encadeada*/

// Estrutura do nó
type Node struct {
	data int
	next *Node
}

// Estrutura da lista encadeada
type LinkedList struct {
	head *Node
}

// Adiciona um elemento no final da lista
func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Remove o primeiro elemento da lista
func (l *LinkedList) RemoveFirst() {
	if l.head != nil {
		l.head = l.head.next
	}
}

// Exibe os elementos da lista
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main2() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("Lista Encadeada Simples:")
	list.Print()
	list.RemoveFirst()
	fmt.Println("Após remover o primeiro:")
	list.Print()
}

/*Exercicio 03*/

// Estrutura do nó (cada tarefa)
type Task struct {
	data string
	next *Task
	prev *Task
}

// Estrutura da lista encadeada (lista de tarefas)
type TaskList struct {
	head *Task
	tail *Task
}

// Método para adicionar uma nova tarefa à lista (Append)
func (t *TaskList) AddTask(data string) {
	newTask := &Task{data: data}
	if t.tail == nil {
		t.head = newTask
		t.tail = newTask
		return
	}
	t.tail.next = newTask
	newTask.prev = t.tail
	t.tail = newTask
}
func (t *TaskList) CompleteTask() {
	if t.tail == nil {
		return
	}
	if t.tail == t.head { // Apenas um elemento na lista
		t.head = nil
		t.tail = nil
		return
	}
	t.tail = t.tail.prev
	t.tail.next = nil
}

func (t *TaskList) ShowTasks() {
	current := t.head
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}
func main3() {
	lista := TaskList{}
	lista.AddTask("Estudar Go")
	lista.AddTask("Fazer compras")
	lista.AddTask("Academia")
	lista.ShowTasks()
	lista.CompleteTask()
	lista.ShowTasks()
	lista.CompleteTask()
	lista.ShowTasks()
}
