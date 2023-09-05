CREATE TYPE requestor AS ENUM (
  'user',
  'sender'
);

CREATE TYPE status AS ENUM (
  'requested',
  'accepted',
  'rejected'
);

CREATE TABLE users (
  id serial PRIMARY KEY,
  email string UNIQUE,
  password string,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

CREATE TABLE senders (
  id serial PRIMARY KEY,
  user_id int,
  device_id string UNIQUE,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

CREATE TABLE user_senders (
  user_id int,
  sender_id int,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  PRIMARY KEY (user_id, sender_id)
);

CREATE TABLE user_sender_reqs (
  id serial PRIMARY KEY,
  user_id int,
  sender_id int,
  requestor requestor,
  status status,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

CREATE TABLE user_api_keys (
  api_key string PRIMARY KEY,
  user int,
  api_secret string,
  expiresAfter timestamp,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

CREATE TABLE batches (
  id serial PRIMARY KEY,
  queued_at timestamp,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now()
);

CREATE TABLE messages (
  id serial PRIMARY KEY,
  to string,
  content string,
  batch_id int,
  sent_at timestamp,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  user_id int,
  sent_from int
);

ALTER TABLE senders ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_senders ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_senders ADD FOREIGN KEY (sender_id) REFERENCES senders (id);
ALTER TABLE user_sender_reqs ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE user_sender_reqs ADD FOREIGN KEY (sender_id) REFERENCES senders (id);
ALTER TABLE user_api_keys ADD FOREIGN KEY (user) REFERENCES users (id);
ALTER TABLE messages ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE messages ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE messages ADD FOREIGN KEY (sent_from) REFERENCES senders (id);
