package models

import "ecommerce-web/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaProdutos() []Produto {
	db := db.ConectBD()

	selectTodosProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectBD()

	insereDadosBd, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosBd.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectBD()

	deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()
}

func BuscaProdutosPorId(id string) Produto {
	db := db.ConectBD()

	produtoSelecionado, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoSelecionado.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoSelecionado.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizar
}

func EditarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectBD()

	editarProduto, err := db.Prepare("UPDATE produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	editarProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
