ALTER TABLE IF EXISTS position_attribute DROP CONSTRAINT fk_position;
ALTER TABLE IF EXISTS position DROP CONSTRAINT fk_profession;
ALTER TABLE IF EXISTS position_attribute DROP CONSTRAINT fk_attribute;
DROP TABLE IF EXISTS position;
DROP TABLE IF EXISTS attribute;
DROP TYPE IF EXISTS states;
DROP TABLE IF EXISTS position_attribute;
DROP TABLE IF EXISTS profession;