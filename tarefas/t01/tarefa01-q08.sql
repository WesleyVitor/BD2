/*Criação da VIEW */
CREATE VIEW nomeDepto_CountProjeto(Gerente,Descricao, Projetos) as SELECT d.cod_gerente,d.descricao,count(p.cod_depto) FROM projeto p, departamento d
WHERE p.cod_depto = d.codigo GROUP BY d.descricao, d.cod_gerente;

/*Seleção */
select f.nome, descricao, projetos from nomeDepto_CountProjeto n LEFT JOIN funcionario f ON f.codigo = n.gerente;