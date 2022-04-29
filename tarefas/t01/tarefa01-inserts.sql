/* Inserindo Departamentos */
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
/* Inserindo Funcionários */
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

/* Inserindo Projetos */
insert into projeto(nome, descricao, cod_depto, cod_responsavel, data_inicio, data_fim)
values ('APF', 'Analisador de Ponto de Função', 2, 19, '2018-02-26', '2019-06-30');

insert into projeto(nome, descricao, cod_depto, cod_responsavel, data_inicio, data_fim)
values ('Monitoria', 'Projeto de Monitoria 2019.1', 4, 21, '2019-02-26', '2019-12-30');

insert into projeto(nome, descricao, cod_depto, cod_responsavel, data_inicio, data_fim)
values ('BD', 'Projeto de Banco de Dados', 3, 22, '2018-02-26', '2018-06-30');

insert into projeto(nome, descricao, cod_depto, cod_responsavel, data_inicio, data_fim)
values ('ES', 'Projeto de Engenharia de Software', 1, 19, '2018-02-26', '2018-06-30');

insert into projeto(nome, descricao, cod_depto, cod_responsavel, data_inicio, data_fim)
values ('DL', 'Projeto de Direito e Legislação', 1, 19, '2019-02-26', '2019-06-30');

/* Inserindo atividades */
insert into atividade(nome, descricao, cod_responsavel, data_inicio, data_fim)
values ('APF - Atividade 1','Fazer atividade 1', 19, '2018-02-26', '2018-06-30');
insert into atividade(nome, descricao, cod_responsavel, data_inicio, data_fim)
values ('APF - Atividade 2','Fazer atividade 2', 19, '2018-02-26', '2018-06-30');
insert into atividade(nome, descricao, cod_responsavel, data_inicio, data_fim)
values ('APF - Atividade 3','Fazer atividade 3', 19, '2018-02-26', '2018-06-30');

insert into atividade(nome, descricao, cod_responsavel, data_inicio, data_fim)
values ('Monitoria - Atividade 1','Fazer atividade 1', 20, '2018-02-26', '2018-06-30');
insert into atividade(nome, descricao, cod_responsavel, data_inicio, data_fim)
values ('Monitoria - Atividade 2','Fazer atividade 1', 20, '2018-02-26', '2018-06-30');

/* Inserindo atividade_projeto */
insert into atividade_projeto(cod_projeto, cod_atividade)
values (1,1);
insert into atividade_projeto(cod_projeto, cod_atividade)
values (2,1);
insert into atividade_projeto(cod_projeto, cod_atividade)
values (3,2);
insert into atividade_projeto(cod_projeto, cod_atividade)
values (1,2);
insert into atividade_projeto(cod_projeto, cod_atividade)
values (4,5);

