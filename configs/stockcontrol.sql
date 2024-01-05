CREATE DATABASE stockcontrol
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Portuguese_Brazil.1252'
    LC_CTYPE = 'Portuguese_Brazil.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
	
CREATE TABLE product_types (
  	id				BIGINT	NOT NULL,
  	description		VARCHAR(200)	NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));
				   
CREATE TABLE products (
  	id          	BIGINT     NOT NULL,
  	description 	VARCHAR(200)     NOT NULL,
  	type_id			BIGINT NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
    PRIMARY KEY ( id ));
				   
CREATE INDEX products_type_id ON products (type_id);
ALTER TABLE products ADD CONSTRAINT products_type_id FOREIGN KEY ( type_id ) REFERENCES product_types(id);


CREATE TABLE labels (
	id          	BIGINT     NOT NULL,
	code 			VARCHAR(2000)     NOT NULL,
	product_id		BIGINT NOT NULL,
	valid_date		DATE,
	created_at		timestamp,
	updated_at		timestamp,
   	PRIMARY KEY ( id ));
				   
CREATE INDEX labels_product_id ON labels(product_id);
ALTER TABLE labels ADD CONSTRAINT labels_product_id FOREIGN KEY ( product_id ) REFERENCES products(id);				   
				   
CREATE TABLE stocks (
  	id				BIGINT	NOT NULL,
  	description		VARCHAR(200)	NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));
				 
CREATE TABLE stock_products (
  	id				BIGINT	NOT NULL,
	stock_id		BIGINT	NOT NULL,
 	product_id		BIGINT	NOT NULL,
	quantity		BIGINT,
	factor			INTEGER,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));

CREATE INDEX stock_products_stock_id ON stock_products (stock_id);
ALTER TABLE stock_products ADD CONSTRAINT stock_products_stock_id FOREIGN KEY ( stock_id ) REFERENCES stocks(id);
CREATE INDEX stock_products_product_id ON stock_products (product_id);
ALTER TABLE stock_products ADD CONSTRAINT stock_products_product_id FOREIGN KEY ( product_id ) REFERENCES products(id);
CREATE UNIQUE INDEX u_stock_products ON stock_products (product_id, stock_id);	 

CREATE TABLE users (
  	id				BIGINT	NOT NULL,
  	name		VARCHAR(200)	NOT NULL,
	email		VARCHAR(100) NOT NULL,
	password	VARCHAR(100) NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));
CREATE UNIQUE INDEX u_user_email ON users(email);
				 
CREATE TABLE user_sessions (
  	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	started_at		timestamp NOT NULL,
	finished_at		timestamp,
	PRIMARY KEY ( id ));
CREATE INDEX user_sessions_user_id ON user_sessions (user_id);
ALTER TABLE user_sessions ADD CONSTRAINT user_sessions_user_id FOREIGN KEY ( user_id ) REFERENCES users(id);
				 
CREATE TABLE operations (
  	id				BIGINT	NOT NULL,
  	name			VARCHAR(200)	NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));		

CREATE TABLE user_operations (
  	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	operation_id 	BIGINT NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id ));	
CREATE INDEX user_operations_user_id ON user_operations (user_id);
ALTER TABLE user_operations ADD CONSTRAINT user_operations_user_id FOREIGN KEY ( user_id ) REFERENCES users(id);
CREATE INDEX user_operations_operation_id ON user_operations (operation_id);
ALTER TABLE user_operations ADD CONSTRAINT user_operations_operation_id FOREIGN KEY ( operation_id ) REFERENCES operations(id);
CREATE UNIQUE INDEX u_user_operations ON user_operations(user_id,operation_id);
				 
CREATE TABLE transactions(
  	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	operation_id 	BIGINT NOT NULL,
	performed_at	timestamp,
	stock_product_id	BIGINT,
	quantity		BIGINT,
	label_id		BIGINT,
	PRIMARY KEY ( id ));	
CREATE INDEX transactions_user_id ON transactions (user_id);
ALTER TABLE transactions ADD CONSTRAINT transactions_user_id FOREIGN KEY ( user_id ) REFERENCES users(id);
CREATE INDEX transactions_operation_id ON transactions (operation_id);
ALTER TABLE transactions ADD CONSTRAINT transactions_operation_id FOREIGN KEY ( operation_id ) REFERENCES operations(id);
CREATE INDEX transactions_label_id ON transactions (label_id);
ALTER TABLE transactions ADD CONSTRAINT transactions_label_id FOREIGN KEY ( label_id ) REFERENCES labels(id);
CREATE INDEX transactions_stock_product_id ON transactions (stock_product_id);
ALTER TABLE transactions ADD CONSTRAINT transactions_stock_product_id FOREIGN KEY ( stock_product_id ) REFERENCES stock_products(id);

				 