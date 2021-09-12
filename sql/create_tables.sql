CREATE table IF NOT EXISTS public.produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade int
);