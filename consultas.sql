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
-- Buscar veiculos cadastrados
select 
	v.id,
	v.placa,
	v.descricao,
	v.categoria_id,
	c.nome as  nome_categoria,
	(
		case
			when v.is_disponivel then 'Disponivel.'
			when v.is_substituido then 'Indisponivel(substituido).'
			else 'Indisponivel.'
		end
	) as status_operacao,
	(
		case
			when v.is_servico then 'Titular.'
			else 'Reserva.'
		end
	) as tipo_veiculo
from veiculo v
left join categoria c on c.id = v.categoria_id
order by v.placa desc
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
	(
	  case 
		when produto.quantidade between 1 and 3 then 'Baixo.'
		when produto.quantidade = 0 then 'Em falta.'
		else 'Normal.'
	  end
	) as status_produto,
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
	produto.quantidade, 
	categoria.id as categoria_id,
	categoria.nome as nome_categoria,
	(
	  case 
		when produto.quantidade between 1 and 3 then 'Baixo.'
		when produto.quantidade = 0 then 'Em falta.'
		else 'Normal.'
	  end
	) as status_produto
from produto
left join categoria on categoria.id = produto.categoria_id
----------------------------------------------------
-- Buscar classificações de manutenções com suas categorias
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
-- Buscar manutenções registradas
select
	m.id,
	m.km_atual,
	m.descricao, 
	m.valor_pago,
	mt.descricao as categoria_manutencao,
	(
		case
			when m.is_finalizado then 'Concluido.'
			when m.is_andamento then 'Em Andamento.'
			when m.is_cancelado then 'Cancelado.'
			else 'Não identificado.'
		end
	) as status_manutencao,
	v.id as id_veiculo_manutencao,
	v.placa as veiculo_manutencao,
	(
		case
			when vs.descricao is null then 'Veiculo não substituido.'
			else vs.placa
		end 
	) as veiculo_substituto,
	vs.id as id_veiculo_substituto,
	u.id as usuario_id,
	u.nome as usuario,
	m.cancelado_em,
	m.created_at as registrado_em
from manutencao m
inner join veiculo v on v.id = m.veiculo_id 
left  join usuario u on u.id = m.usuario_id 
inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
left join veiculo vs on vs.id = m.veiculo_id_temporario
order by m.id desc;
----------------------------------------------------
-- Consultas previas detalhadas
select 
	mp.id,
	mp.km_previsao,
	mp.data_previsao,
	mt.descricao as categoria_manutencao,
	(
		case
			when mp.is_aprovado then 'Aprovado.'
			else 'Cancelado.'
		end
	) as status_previsao,
	m.km_atual,
	m.descricao, 
	m.valor_pago,
	mt.descricao as categoria_manutencao,
	(
		case
			when m.is_finalizado then 'Concluido.'
			when m.is_andamento then 'Em Andamento.'
			when m.is_cancelado then 'Cancelado.'
			else 'Não identificado.'
		end
	) as status_manutencao,
	v.id as id_veiculo_manutencao,
	v.placa as veiculo_manutencao,
	(
		case
			when vs.descricao is null then 'Veiculo não substituido.'
			else vs.placa
		end 
	) as veiculo_substituto,
	vs.id as id_veiculo_substituto,
	u.id as usuario_id,
	u.nome as usuario,
	m.created_at as manutencao_realizado,
	mp.created_at as previsao_aberta_em,
	mp.updated_at as ultima_atualizacao_registro
FROM manutencao_previsao mp
inner join manutencao m on m.id = mp.ultima_manutencao_id 
inner join veiculo v on v.id = m.veiculo_id 
left  join usuario u on u.id = m.usuario_id 
inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
left join veiculo vs on vs.id = m.veiculo_id_temporario
where mp.ultima_manutencao_id = 1;
----------------------------------------------------
-- Listar todas previsões do dia
select 
	mp.id,
	mp.km_previsao,
	mp.data_previsao,
	mt.descricao as categoria_manutencao,
	(
		case
			when mp.is_aprovado then 'Aprovado.'
			else 'Cancelado.'
		end
	) as status_previsao,
	m.descricao, 
	v.placa as veiculo_manutencao,
	mp.created_at as previsao_aberta_em,
	mp.updated_at as ultima_atualizacao_registro
FROM manutencao_previsao mp
inner join manutencao m on m.id = mp.ultima_manutencao_id 
inner join veiculo v on v.id = m.veiculo_id 
inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
where date(mp.created_at ) = current_date
and m.is_cancelado = false
and mp.is_cancelado = false
order by mp.created_at desc;
-- Listar todas previsões
select 
	mp.id,
	mp.km_previsao,
	mp.data_previsao,
	mt.descricao as categoria_manutencao,
	(
		case
			when mp.is_aprovado then 'Aprovado.'
			else 'Cancelado.'
		end
	) as status_previsao,
	m.descricao, 
	v.placa as veiculo_manutencao,
	mp.created_at as previsao_aberta_em,
	mp.updated_at as ultima_atualizacao_registro
FROM manutencao_previsao mp
inner join manutencao m on m.id = mp.ultima_manutencao_id 
inner join veiculo v on v.id = m.veiculo_id 
inner join manutencao_tipo mt on mt.id = m.manutencao_tipo_id 
order by mp.created_at desc;
----------------------------------------------------
-- Consultas básicas
select * from usuario;
select * from veiculo;
select * from categoria;
select * from produto;
select * from movimento_produto;
select * from manutencao;
select * from manutencao_tipo;
select * from manutencao_previsao;
----------------------------------------------------