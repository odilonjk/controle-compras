CREATE TYPE enum_satisfacao AS enum('muito_satisfeito', 'satisfeito', 'insatisfeito');

CREATE TYPE enum_pagamento AS enum('dinheiro', 'cartao', 'boleto');

CREATE TABLE compras (
	id serial PRIMARY KEY,
	valor numeric NOT NULL DEFAULT 0,
	"data" date NOT NULL,
	observacao varchar(255) NOT NULL,
	recebido int2,
	forma_pagamento enum_pagamento NOT NULL,
	satisfacao enum_satisfacao NULL
);