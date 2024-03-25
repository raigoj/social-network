package structs

import "time"

type User struct {
	Id        int
	Username  string
	Age       string
	Gender    string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Online    int
}

type Chat struct {
	User                User
	LastMessage         Message
	UnreadMessagesCount int
}

type Session struct {
	Sessionid    string
	Userid       int
	Lastactivity time.Time
}

type Posts struct {
	Id           int
	Title        string
	Content      string
	User         int
	Category     int
	Creationtime string
	Username     string
}

type Comment struct {
	Id           int
	Content      string
	Username     string
	Postid       int
	Creationtime string
}

type Category struct {
	Id   int
	Name string
}

type Message struct {
	Id         int
	Senderid   int
	Receiverid int
	Text       string
	Sentat     string
	Read       int
}

type Group struct {
	Id      int
	Creator int
	Name    string
	Text    string
}

type UserGoup struct {
	Uid int
	Gid int
}

type Event struct {
	Id      int
	Creator int
	Name    string
	Text    string
}

type UserEvent struct {
	Uid    int
	Eid    int
	Status string
}

type Invite struct {
	Sid int
	Rid int
	Gid int
}
