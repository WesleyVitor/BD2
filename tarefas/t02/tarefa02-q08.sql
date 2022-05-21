CREATE FUNCTION delete_cliente()
RETURNS trigger AS $$
BEGIN
	DELETE FROM cliente where codigo = new.cliente;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER delete_cliente_by_voo 
AFTER DELETE ON cliente_voo
    FOR EACH ROW 
    EXECUTE FUNCTION delete_cliente();
