SELECT f.nome, f.salario, d.descricao
FROM funcionario f, departamento d
WHERE f.codDepto = d.codigo;