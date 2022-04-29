drop table if exists atividade cascade;
drop table if exists projeto cascade;
drop table if exists departamento cascade;
drop table if exists funcionario cascade;
drop table if exists atividade_projeto cascade;

CREATE TABLE funcionario (
	codigo SERIAL PRIMARY KEY,
	nome varchar(50),
	sexo char(1),
	dtNasc date,
	salario decimal(10,2),
	codDepto int
);

CREATE TABLE departamento (
	codigo SERIAL PRIMARY KEY,
	descricao varchar(50),
	cod_gerente int,
	FOREIGN KEY (cod_gerente) REFERENCES funcionario(codigo) on delete set null on update cascade
);

ALTER TABLE funcionario ADD FOREIGN KEY (codDepto) REFERENCES departamento(codigo) on delete set null on update cascade;


CREATE TABLE projeto (
	codigo SERIAL PRIMARY KEY,
	nome varchar(50),
	descricao varchar(250),
	cod_responsavel int,
	cod_depto int,
	data_inicio date, 
	data_fim date,
	FOREIGN KEY (cod_responsavel) REFERENCES funcionario(codigo) on delete set null on update cascade,
	FOREIGN KEY (cod_depto) REFERENCES departamento(codigo) on delete set null on update cascade
);

CREATE TABLE atividade (
	codigo SERIAL PRIMARY KEY,
	nome varchar(50),
	descricao varchar(250),
	cod_responsavel int,
	data_inicio date, 
	data_fim date,
	FOREIGN KEY (cod_responsavel) REFERENCES funcionario(codigo) on delete set null on update cascade

);

CREATE TABLE atividade_projeto(
  cod_projeto INT NOT NULL, 
  cod_atividade INT NOT NULL,

  PRIMARY KEY(cod_projeto, cod_atividade),
  FOREIGN KEY(cod_projeto) REFERENCES projeto(codigo) on delete set null on update cascade,
  FOREIGN KEY(cod_atividade) REFERENCES atividade(codigo) on delete set null on update cascade
)