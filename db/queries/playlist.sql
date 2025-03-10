-- name: CreatePlaylist :one
INSERT INTO playlists (
  user_id,
  spotify_id,
  name
) VALUES ( $1, $2, $3 )
RETURNING *;

-- name: GetPlaylistsByUserID :many
SELECT * FROM playlists
WHERE user_id = $1;
