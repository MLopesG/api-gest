----------------------------------------------------
-- Buscar manutenções registradas
select
	manutencao.id, 
	veiculo.placa, 
	usuario.nome as usuario, 
	manutencao.km_atual, 
	manutencao.descricao, 
	categoria.nome as nome_categoria
from manutencao
left join categoria on categoria.id = manutencao.categoria_id
left join usuario on usuario.id = manutencao.usuario_id
left join veiculo on veiculo.id = manutencao.veiculo_id
order by manutencao.id desc;
----------------------------------------------------
-- Buscar Movimentações do(s) veiculo(s) registrado(s)
select
	movimento_veiculo.id, 
	veiculo.placa, 
	usuario.nome as usuario,
	movimento_veiculo.data_saida_entrada,
	movimento_veiculo.data_retorno_chegada,
	movimento_veiculo.tipo_movimento,
	movimento_veiculo.km_saida,
	movimento_veiculo.km_entrada,
	movimento_veiculo.created_at as registrado_em
from movimento_veiculo
left join usuario on usuario.id = movimento_veiculo.usuario_id
left join veiculo on veiculo.id = movimento_veiculo.veiculo_id
order by movimento_veiculo.id desc;
----------------------------------------------------
-- Buscar histórico de saidas e entradas dos produtos
select 
	movimento_produto.id,
	movimento_produto.destino_produto,
	produto.nome as produto,
	usuario.nome as usuario,
	veiculo.placa as veiculo_placa,
	(
	  case 
		when movimento_produto.is_entrada then 'Entrada de Produto.'
		else 'Saida de Produto.'
	  end
	) as tipo_movimento,
	movimento_produto.quantidade,
	produto.quantidade as quantidade_em_estoque,
	movimento_produto.created_at as registrado_em
from  movimento_produto
inner join usuario on usuario.id = movimento_produto.usuario_id
inner join produto on produto.id = movimento_produto.produto_id
left join veiculo on veiculo.id = movimento_produto.veiculo_id
order by movimento_produto.created_at desc;
----------------------------------------------------
-- Buscar Produtos e suas categorias relacionadas
select 
	produto.id, 
	produto.nome, 
	categoria.id as categoria_id,
	categoria.nome as nome_categoria
 from produto
 left join categoria on categoria.id = produto.categoria_id;
----------------------------------------------------
-- Consultas básicas
select * from usuario;
select * from veiculo;
select * from categoria;
select * from produto;
select * from movimento_produto;
select * from manutencao;
----------------------------------------------------
-- Buscar calssificações de manutenções com suas categorias
select 
	t.id,
	t.descricao,
	t.km_previsto, 
	t.intervalo_previsto,
	c.id,
	c.nome as categoria
from manutencao_tipo t
inner join categoria c on c.id = t.categoria_id 
order by t.descricao;
----------------------------------------------------