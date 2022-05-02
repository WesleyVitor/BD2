SELECT p.nome, p.data_inicio, p.data_fim, a.descricao, a.data_inicio, a.data_fim
FROM atividade a, projeto p, atividade_projeto ap
WHERE a.codigo = ap.cod_atividade and p.codigo = ap.cod_projeto;