-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "unaccent";
CREATE EXTENSION IF NOT EXISTS "hstore";

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
DECLARE
    col TEXT := quote_ident(TG_ARGV[0]);
BEGIN
    NEW := NEW #= hstore(col, NOW()::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION slugify("value" TEXT)
RETURNS TEXT AS $$
  -- removes accents (diacritic signs) from a given string --
  WITH "unaccented" AS (
    SELECT unaccent("value") AS "value"
  ),
  -- lowercases the string
  "lowercase" AS (
    SELECT lower("value") AS "value"
    FROM "unaccented"
  ),
  -- replaces anything that's not a letter, number, "-", or underscore('_') with "-"
  "hyphenated" AS (
    SELECT regexp_replace("value", '[^a-z0-9\\_-]+', '-', 'gi') AS "value"
    FROM "lowercase"
  ),
  -- trims "-" if they exist on the head or tail of the string
  "trimmed" AS (
    SELECT regexp_replace(regexp_replace("value", '-\\+$', ''), '^\\-', '') AS "value"
    FROM "hyphenated"
  )
    SELECT "value" FROM "trimmed";
$$ LANGUAGE SQL STRICT IMMUTABLE;

-- creates a slug from the src col into the dst column using the specified separator
CREATE FUNCTION set_slug_from_col()
 RETURNS TRIGGER
AS $$
DECLARE
  src TEXT := quote_ident(TG_ARGV[0]);
 	dst  TEXT := quote_ident(TG_ARGV[1]);
 	slug TEXT;
BEGIN
	EXECUTE 'SELECT $1.' || src 
	USING NEW
	INTO src; 
 
	slug := reno.slugify(src);
	
	NEW := NEW #= hstore(dst, slug);
 
  RETURN NEW;
END
$$ LANGUAGE plpgsql;

-- +migrate StatementEnd

-- +migrate Down
-- SQL in section 'Up' is executed when this migration is applied