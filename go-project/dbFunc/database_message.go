package dbFunc

import (
	"fmt"
	"time"

	"social-network/structs"
)

func SetMessage(dbName string, senderid int, receiverid int, text string) error {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	creationtime := time.Now()
	insertSQL := `INSERT INTO MESSAGE(senderid, receiverid, text, sentat, read) VALUES (?, ?, ?, ?, ?)`
	statement, err := forumDB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(senderid, receiverid, text, creationtime, 0)
	if err != nil {
		return err
	}
	return nil
}

/*
This SQL query is used to retrieve a list of messages from a messages table. The query selects the id, sender_id, recipient_id, message, date, and read columns from the table. It filters the results using the following conditions:

	The sender_id must be equal to $1 and the recipient_id must be equal to $2, or vice versa. This ensures that only messages sent between the specified sender and recipient are returned.

	The id of the message must be less than $3 if $3 is not equal to 0, otherwise any value is allowed. This condition is used to retrieve only messages that were sent before a specified message ID.

The query sorts the results by the id column in descending order and limits the number of results to $4.

In the parameters for the query, $1, $2, $3, and $4 are placeholders for the values of senderID, recipientID, lastMessageID, and limit, respectively. These values are passed to the query when it is executed to retrieve the desired results.

Overall, this query is used to retrieve a list of messages sent between two specified users, possibly limited to messages that were sent before a specified message ID and limited to a specified number of results.
*/
func UpdateMessageRead(dbName string, receiverid int, senderid int) error {
	forumDB := OpenDatabase((dbName))
	defer forumDB.Close()
	_, err := forumDB.Exec("UPDATE message SET read = 1 WHERE senderid = ? AND receiverid = ?", senderid, receiverid)
	if err != nil {
		fmt.Println("error krt")
		return err
	}
	return nil
}

func GetMessages(dbName string, senderid int, receiverid int, limit int) ([]structs.Message, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	var message structs.Message
	messages := make([]structs.Message, 0)
	// CHATGPT ftw
	offset := limit
	limit = 10
	rows, err := forumDB.Query(`SELECT * FROM (
		SELECT id, text, sentat, senderid, receiverid, read
		FROM message
		WHERE (senderid = $1 AND receiverid = $2) OR (receiverid = $1 AND senderid = $2)
		ORDER BY id DESC
		LIMIT $3 OFFSET $4
	  ) as messages
	  ORDER BY id ASC`,
		senderid, receiverid, limit, offset)
	if offset != 0 {
		rows, err = forumDB.Query(`
			SELECT id, text, sentat, senderid, receiverid, read
			FROM message
			WHERE (senderid = $1 AND receiverid = $2) OR (receiverid = $1 AND senderid = $2)
			ORDER BY id DESC
			LIMIT $3 OFFSET $4
		  `,
			senderid, receiverid, limit, offset)
		//forumDB.Query("SELECT id, text, sentat, senderid, receiverid FROM message WHERE senderid = ? AND receiverid = ? OR receiverid = ? AND senderid = ? ORDER BY id DESC LIMIT 10;", senderid, receiverid, senderid, receiverid)
	}
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&message.Id, &message.Text, &message.Sentat, &message.Senderid, &message.Receiverid, &message.Read); err {
		case nil:
			messages = append(messages, message)
		default:
			return messages, err
		}
	}
	return messages, nil
}

func GetUsersWithLastMessage(dbName string, userid int) ([]structs.Chat, error) {
	forumDB := OpenDatabase(dbName)
	defer forumDB.Close()

	var chat structs.Chat
	chats := make([]structs.Chat, 0)
	rows, err := forumDB.Query(`
		SELECT 
			users.id, 
			users.firstname, 
			users.lastname,
			IFNULL(MAX(m.id), 0), 
			IFNULL(m.senderid, 0), 
			IFNULL(m.receiverid, 0), 
			IFNULL(m.text, 0), 
			IFNULL(m.sentat, 0),
			SUM(CASE WHEN (m.read = 0 AND m.receiverid = $1) THEN 1 ELSE 0 END) AS unread_messages_count 
		FROM users 
		LEFT JOIN message m ON 
			m.senderid = $1 AND m.receiverid = users.id
		OR 
			m.senderid = users.id AND m.receiverid = $1 
		WHERE NOT users.id = $1 
		GROUP BY users.id 
		ORDER BY 
			m.id DESC, 
			users.firstname ASC
	`,
		userid)
	if err != nil {
		return chats, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&chat.User.Id, &chat.User.Firstname, &chat.User.Lastname, &chat.LastMessage.Id, &chat.LastMessage.Senderid, &chat.LastMessage.Receiverid, &chat.LastMessage.Text, &chat.LastMessage.Sentat, &chat.UnreadMessagesCount); err {
		case nil:
			fmt.Println("unreadmessagecount = ", chat.UnreadMessagesCount)
			chats = append(chats, chat)
		default:
			return chats, err
		}
	}
	return chats, nil
}
