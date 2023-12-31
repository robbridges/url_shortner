CREATE TABLE IF NOT EXISTS url (
  id SERIAL PRIMARY KEY,
  url VARCHAR(255) NOT NULL,
  short_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  UNIQUE (url)
);