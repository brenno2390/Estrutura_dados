package main

import (
	"fmt"
)

// Definição da pilha
type Stack2 struct {
	items []rune
}

// Push adiciona um caractere na pilha
func (s *Stack2) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o caractere do topo da pilha
func (s *Stack2) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false // Retorna falso se a pilha estiver vazia
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Remove o último elemento
	return item, true

}

// Função para verificar se uma palavra é um palíndromo usando pilha
func isPalindrome(word string) bool {
	s := Stack2{}
	// Adiciona cada caractere da palavra na pilha
	for _, char := range word {
		s.Push(char)
	}

	// Verifica se os caracteres retirados da pilha são iguais aos da palavra original
	for _, char := range word {
		poppedChar, _ := s.Pop()
		if poppedChar != char {
			return false
		}
	}
	return true
}

func main2() {
	// Testes
	words := []string{"radar", "arara", "golang", "level", "hello"}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}
