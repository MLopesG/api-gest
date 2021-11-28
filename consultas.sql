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

select 
	produto.id, 
	produto.nome, 
	categoria.id as categoria_id,
	categoria.nome as nome_categoria
 from produto
 left join categoria on categoria.id = produto.categoria_id;

select * from usuario;
select * from veiculo;
select * from categoria;
select * from manutencao;
select * from produto;
select * from movimento_produto;

