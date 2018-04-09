package main

type Player struct {
	Name  string `valid:"alphanum, length(0|10)"`
	Email string `valid:"email"`
}

type Message struct {
	Message     string `valid:"ascii"`
	Sender      string `valid:"ascii"`
	Receiver    string `valid:"ascii"`
	MessageType string `valid:"messagetype~Message Type: messagetype must be text or"`
	Read        bool   `valid:"read"`
}

type Challenge struct {
	Title           string `valid:"ascii"`
	Description     string `valid:"ascii"`
	CompetitionMode string `valid:"competitionmode~Competition Mode: Competitionmode must be pvp or"`
	PlayCategory    string `valid:"-"`
	Target          string `valid:"-"`
	MaxPlayer       int64  `valid:"numberofplayer~Max Player: must be a number between 0 and 1000"`
	Start           int64  `valid:"challengedate"`
	End             int64  `valid:"challengedate"`
	Completed       int64  `valid:"challengedate"`
	Image           string `valid:"url"`
	Master          string `valid: "ascii"`
}
