package main

import "net/http"

// Go possui tipagem forte, minha string vai ser sempre string
// essa estrutura abaixo tem o nome de 'struc', é uma estrutura de dados que vc consegue colocar varios tipos dentro dela
// na ultima coluna, eu consigo colocar como os dados vao aparecer em uma aventual exportação para json
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

// Esse debaixo é um slice que é um "array dinamico", Go tambem possui array's, porém eles sao fixos, slices sao dinamicos
var Products = []Product{
	{
		ID:          1,
		Name:        "Smartphone Galaxy S23",
		Description: "Flagship da Samsung com câmera de 108MP",
		Price:       4599.90,
		Category:    "Eletrônicos",
	},
	{
		ID:          2,
		Name:        "Notebook Dell XPS 13",
		Description: "Ultrabook com tela InfinityEdge",
		Price:       7899.00,
		Category:    "Informática",
	},
	{
		ID:          3,
		Name:        "Camiseta Básica Algodão",
		Description: "Camiseta 100% algodão orgânico",
		Price:       79.90,
		Category:    "Vestuário",
	},
	{
		ID:          4,
		Name:        "Fone Bluetooth Sony WH-1000XM5",
		Description: "Fone com cancelamento de ruído premium",
		Price:       2199.99,
		Category:    "Áudio",
	},
	{
		ID:          5,
		Name:        "Livro 'Clean Code'",
		Description: "Robert Martin's guide to better coding practices",
		Price:       129.50,
		Category:    "Livros",
	},
}

func main() {
	// println("oie")
	http.ListenAndServe(":8080", nil)
}
