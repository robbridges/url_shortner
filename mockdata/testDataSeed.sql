CREATE TABLE IF NOT EXISTS url (
   id SERIAL PRIMARY KEY,
   url VARCHAR(255) NOT NULL,
   short_url VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   leetcode boolean DEFAULT false NOT NULL,
   UNIQUE (url)
);

INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://www.google.com', 'abcdef', NOW(), NOW(), false);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem', 'rndpr1', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem2', 'rndpr2', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem3', 'rndpr3', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem4', 'rndpr4', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem5', 'rndpr5', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem6', 'rndpr6', NOW(), NOW(), true);
INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ('https://leetcode.com/randomproblem7', 'rndpr7', NOW(), NOW(), true);