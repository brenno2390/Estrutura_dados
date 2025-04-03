// 1 e 2
package main

import "fmt"

func main() {
	var arr [6]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	arr[3] = 6
	arr[4] = 2
	arr[5] = 6

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	fmt.Println(arr[4])
	fmt.Println(arr[5])
}

//3-

func main3() {
	var arr [6]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	arr[3] = 6
	arr[4] = 2
	arr[5] = 6

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

//4-

func main4() {
	var arr [6]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 4
	arr[3] = 6
	arr[4] = 2
	arr[5] = 6

	for i, v := range arr {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}
}

//5-

func main5() {
	var matriz [3][3]string
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("Digite o valor do índice [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])

		}
	}
}

//6-

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func main6() {
	produto := Produto{
		Nome:       "Notebook",
		Preco:      3599.99,
		Quantidade: 50,
	}

	// Acessando os campos do produto
	fmt.Println("Nome do produto:", produto.Nome)
	fmt.Printf("Preço: R$ %.2f\n", produto.Preco)
	fmt.Println("Quantidade em estoque:", produto.Quantidade)
}
