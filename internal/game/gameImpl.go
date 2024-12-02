package game

// Implementation of main game logic
// Server sends 'init' request -> game beginning
// Every 2 seconds need to send {Score} {BatteryStatus} to server
//
// If route is '/' -> server must run new session
// If route is '/:id' where id is Telegram ID written in requirements
// Server loads client's configuration. and continue the game (sends init message)
