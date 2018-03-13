create table accounts (
  id     serial primary key,
  name   text,
  email  varchar(255)
);

INSERT INTO accounts (name, email) VALUES ('testuser', 'testuser@test.com')