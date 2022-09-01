CREATE TABLE products (
  id BIGSERIAL PRIMARY KEY,
  is_active boolean DEFAULT true,
  product_type integer,
  name character varying,
  description text,
  sku character varying(25),
  price numeric(15,2) DEFAULT 0.0,
  search_vector tsvector,
  destroyed_at timestamp(0) without time zone,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) without time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX products_search_vector_index ON products USING gin (search_vector);
CREATE INDEX products_destroyed_at_index ON products USING btree (destroyed_at);

CREATE OR REPLACE FUNCTION full_search_for_products() RETURNS trigger AS $$
begin
  new.search_vector :=
      setweight(to_tsvector('pg_catalog.simple', coalesce(new.name, '')), 'A') ||
      setweight(to_tsvector('pg_catalog.simple', coalesce(new.sku, '')), 'B');
    return new;
  end
$$ LANGUAGE plpgsql;

CREATE TRIGGER product_search BEFORE INSERT OR UPDATE 
ON products 
FOR EACH ROW EXECUTE PROCEDURE full_search_for_products();