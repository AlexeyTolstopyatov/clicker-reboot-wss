package game

// 	 When client attempts to connect
// Server must detect :id with array's segment in PostgresSQL
// If all ok, Server knows all cells (id, tid, team, score)
//   When unknown client attempts to connect
// Server must deretmine team of unregirtered user, and send initialization request
//

// GlobalScore
// Struct that updates in DB every [xxx] seconds
// Writes or Reads from/to database when server closing/opening
// Use ShutdownService() to write unsaved changes
// and InitializeService() to read all saved information from database.
type GlobalScore struct {
	WhiteTeamClicks   int64
	BlackTeamClicks   int64
	WhiteClickPercent float64
	BlackClickPercent float64
}