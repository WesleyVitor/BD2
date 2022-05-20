CREATE FUNCTION soma_milhas()
RETURNS trigger AS $$
BEGIN
	UPDATE milhas SET quantidade = quantidade + 
	(SELECT distancia FROM voo, cliente_voo 
	WHERE new.voo = voo.codigo
	) / 10
	WHERE  cliente = new.cliente;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_client
AFTER INSERT ON cliente_voo
    FOR EACH ROW 
    EXECUTE FUNCTION soma_milhas();
