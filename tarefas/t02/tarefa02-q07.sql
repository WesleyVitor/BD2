CREATE FUNCTION add_passageiro()
RETURNS trigger AS $$
BEGIN
	UPDATE voo SET num_passageiros = num_passageiros + 1;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER add_milhas 
AFTER INSERT ON cliente_voo
    FOR EACH ROW 
    EXECUTE FUNCTION add_passageiro();
