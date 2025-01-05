CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  username VARCHAR(20) UNIQUE NOT NULL,
  first_name VARCHAR(30),
  last_name VARCHAR(30),
  role VARCHAR(20) DEFAULT 'user',
  encrypted_password VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME
);
