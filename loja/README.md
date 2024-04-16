```shell
docker run --rm --name loja-alura-psql -p 5432:5432 -v /tmp/database:/var/lib/postgresql/data -e POSTGRES_PASSWORD=1234 -d postgres
```

```sql
create database alura_loja ;

create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
) ;

insert into produtos (nome, descricao, preco, quantidade) values
('Camiseta', 'Hugo Boss', 99.9, 9),
('Boné', 'NBA Oficial', 149.9, 7),
('Laptop', 'Apple', 2499.9, 5) ;
```
