SELECT nome, salario, codDepto
FROM funcionario
WHERE codigo NOT IN ( SELECT cod_gerente FROM departamento);