CREATE VIEW mostraCount(codigoF, quantidadeA) AS SELECT f.codigo, count(a.codigo) 
FROM atividade a, funcionario f
WHERE a.cod_responsavel = f.codigo
GROUP BY f.codigo;

SELECT f.nome, mc.codigoF, mc.quantidadeA 
FROM funcionario f, mostraCount mc;