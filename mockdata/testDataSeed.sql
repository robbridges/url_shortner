CREATE TABLE IF NOT EXISTS url (
   id SERIAL PRIMARY KEY,
   url VARCHAR(255) NOT NULL,
   short_url VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   leetcode boolean DEFAULT false NOT NULL,
   UNIQUE (url)
);

INSERT INTO url (url, short_url, created_at, updated_at) VALUES ('https://www.google.com', 'abcdef', NOW(), NOW());