CREATE FUNCTION criar_milhagem()
RETURNS trigger AS $$
BEGIN
	INSERT INTO milhas(cliente, quantidade) 
  	values (NEW.codigo, 0);
  	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER criar_milha
AFTER INSERT ON cliente
    FOR EACH ROW 
    EXECUTE FUNCTION criar_milhagem();
