package models

import "github.com/thukabjj/go-crud-mvc/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConnectComBanco()
	defer db.Close()
	selectTodosOsProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		p := Produto{}
		err = selectTodosOsProdutos.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, p)
	}
	return produtos
}

func CriarNovoProduto(p *Produto) {
	db := db.ConnectComBanco()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	validateError(err)

	insereDadosNoBanco.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
}
func DeletaProduto(idProduto int) {
	db := db.ConnectComBanco()
	defer db.Close()
	deletaDadosNoBanco, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	validateError(err)
	deletaDadosNoBanco.Exec(idProduto)
}

func BuscarProdutoPorId(idProduto int) Produto {
	db := db.ConnectComBanco()
	defer db.Close()

	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", idProduto)
	validateError(err)

	p := Produto{}

	for produtoDoBanco.Next() {
		err := produtoDoBanco.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		validateError(err)
	}
	return p
}
func AtualizarProduto(p *Produto) {
	db := db.ConnectComBanco()
	defer db.Close()
	atualizaDadosNoBanco, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	validateError(err)
	atualizaDadosNoBanco.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)
}
func validateError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
