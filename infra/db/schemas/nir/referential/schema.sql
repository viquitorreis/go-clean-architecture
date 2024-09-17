CREATE SCHEMA nir;

CREATE TABLE nir.referential (
	referential_id bigserial NOT NULL,
	"name" varchar(255) NOT NULL,
	"label" varchar(255) NOT NULL,
	image_path text NULL,
	params jsonb NULL,
	company_id int4 NULL,
	user_audit_id varchar(255) NULL,
	CONSTRAINT referential_pkey PRIMARY KEY (referential_id)
);