CREATE SCHEMA nir;

CREATE TABLE nir."result" (
	result_id bigserial NOT NULL,
	sample_name varchar(255) NOT NULL,
	spectro text NOT NULL,
	company_id int4 NULL,
	operation_id int4 NOT NULL,
	product_id int4 NOT NULL,
	user_audit_id int4 NULL,
	observation text NULL,
	operator_id int4 NULL,
	user_id int4 NULL,
	CONSTRAINT result_pkey PRIMARY KEY (result_id),
	CONSTRAINT result_ukey_prdc UNIQUE (product_id, operation_id)
);
