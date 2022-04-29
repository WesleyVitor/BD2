insert into departamento
(descricao, cod_gerente)
values ( 'Dep História', null);

insert into departamento
(descricao, cod_gerente)
values ( 'Dep Computação', null);

insert into departamento
( descricao, cod_gerente)
values ( 'Dep Geografia', null);

insert into departamento
( descricao, cod_gerente)
values (null, null);

insert into departamento
( descricao, cod_gerente)
values ('Dep Direito', null);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('Ana', 'F', '1988-05-07', 2500.00, 1);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('Taciano', 'M', '1980-01-25', 2500.00, 2);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('Maria', 'F', '1980-01-25', 3500.00, 3);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('João', 'M', '1990-01-25', 1500.00, 3);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('Bianca', 'F', '1960-01-25', 7500.00, 4);

insert into funcionario
(nome, sexo, dtNasc, salario, codDepto)
values ('Bernardo', 'M', '1995-01-25', 3500.00, 5);

update departamento set cod_gerente = 19 where codigo = 2;
update departamento set cod_gerente = 20 where codigo = 3;
update departamento set cod_gerente = 22 where codigo = 4;
update departamento set cod_gerente = 23 where codigo = 5;