CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  display_name VARCHAR(255),
  email VARCHAR(255),
  spotify_id VARCHAR(255) NOT NULL,
  access_token TEXT NOT NULL,
  refresh_token TEXT NOT NULL
);

CREATE TABLE playlists (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  spotify_id TEXT NOT NULL,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE playlist_artists (
  id SERIAL PRIMARY KEY,
  playlist_id INTEGER NOT NULL REFERENCES playlists(id),
  spotify_id TEXT NOT NULL,
  name VARCHAR(255) NOT NULL
);
