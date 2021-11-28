
CREATE TABLE usuario (
  id serial PRIMARY KEY,
  nome VARCHAR(255),
  is_ativo boolean default true,
  senha VARCHAR(255),
  cpf VARCHAR(255),
  email VARCHAR(255),
  created_at TIMESTAMP default now(),
  updated_at TIMESTAMP default now()
  
);

CREATE TABLE categoria(
	id serial PRIMARY KEY,
	nome VARCHAR(100),
	is_categoria_produto boolean default false,
	is_categoria_manutencao boolean default false,
	is_categoria_veiculo boolean default false,
	created_at timestamp default now(),
	updated_at timestamp default now()
);

CREATE TABLE veiculo (
	id serial PRIMARY KEY,
	placa VARCHAR(255),
	descricao VARCHAR(255),	
	categoria_id INTEGER,
	is_servico boolean default true,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(categoria_id) REFERENCES categoria (id)
);

CREATE TABLE manutencao (
	id serial PRIMARY KEY,
	km_atual INTEGER,
	descricao text,
	veiculo_id INTEGER,
	categoria_id INTEGER,
	usuario_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(veiculo_id) REFERENCES veiculo (id),
	FOREIGN KEY(categoria_id) REFERENCES categoria (id),
	FOREIGN KEY(usuario_id) REFERENCES usuario (id)
);

CREATE TABLE movimento_veiculo (
	id serial PRIMARY KEY,
	data_saida_entrada timestamp,
	data_retorno_chegada timestamp,
	km_saida INTEGER,
	km_entrada INTEGER,
	usuario_id INTEGER,
	tipo_movimento VARCHAR(255),
	veiculo_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(veiculo_id) REFERENCES veiculo (id),
	FOREIGN KEY(usuario_id) REFERENCES usuario (id)
);

CREATE TABLE produto (
	id serial PRIMARY KEY,
	nome VARCHAR(100),
	data_entrada_estoque timestamp,
	categoria_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(categoria_id) REFERENCES categoria (id)
);

CREATE TABLE movimento_produto (
	id serial PRIMARY KEY,
	destino_produto VARCHAR(255),
	quantidade INTEGER,
	produto_id INTEGER,
	usuario_id INTEGER,
	veiculo_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(produto_id) REFERENCES produto (id),
	FOREIGN KEY(usuario_id) REFERENCES usuario (id)
);