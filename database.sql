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
	is_categoria_tipo_manutencao boolean default false,
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
	is_reserva boolean default false,
	is_disponivel boolean default true,
	is_indisponivel boolean default false,
	is_substituido boolean default false,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(categoria_id) REFERENCES categoria (id)
);

CREATE TABLE tipo_manutencao(
	id serial PRIMARY KEY,
	descricao VARCHAR(255),
	km_previsto INTEGER default 0,
	intervalo_previsto INTEGER,
	categoria_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(categoria_id) REFERENCES categoria (id)
);

CREATE TABLE manutencao (
	id serial PRIMARY KEY,
	km_atual INTEGER,
	descricao text,
	veiculo_id INTEGER,
	tipo_manutencao_id INTEGER,
	usuario_id INTEGER,
	is_finalizado boolean default false,
	veiculo_id_temporario INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(veiculo_id) REFERENCES veiculo (id),
	FOREIGN KEY(tipo_manutencao_id) REFERENCES tipo_manutencao (id),
	FOREIGN KEY(usuario_id) REFERENCES usuario (id)
);

CREATE TABLE manutencao_previsao (
	id serial PRIMARY KEY,
	km_previsao INTEGER,
	is_finalizado boolean default false,
	data_previsao date,
	manutencao_id INTEGER,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(veiculo_id) REFERENCES veiculo (id),
	FOREIGN KEY(manutencao_id) REFERENCES manutencao (id)
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
	quantidade INTEGER default 0,
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
	is_entrada boolean default false,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	FOREIGN KEY(produto_id) REFERENCES produto (id),
	FOREIGN KEY(usuario_id) REFERENCES usuario (id)
);