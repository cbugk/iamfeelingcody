CREATE TABLE IF NOT EXISTS GithubUsers (
  id   INTEGER PRIMARY KEY AUTOINCREMENT,
  timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
  name TEXT   UNIQUE NOT NULL,
  ralpv INTEGER UNIQUE NOT NULL,
  present BOOLEAN NOT NULL
);
