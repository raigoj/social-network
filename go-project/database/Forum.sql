CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "email" varchar UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL, 
  "gender" varchar NOT NULL,
  "dateofbirth" DATE,
  "online" INTEGER
);

CREATE TABLE IF NOT EXISTS "posts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "user" int,
  "category" int,
  "creationtime" DATE,
  "username" varchar NOT NULL,
  FOREIGN KEY (user)
      REFERENCES users (id),
  FOREIGN KEY (category)
      REFERENCES categories (id)
);

CREATE TABLE IF NOT EXISTS "comments" (
  "id" INTEGER PRIMARy KEY AUTOINCREMENT,
  "content" varchar NOT NULL,
  "username" varchar NOT NULL,
  "postid" int,
  "creationtime" DATE,
  FOREIGN KEY (username)
      REFERENCES users (username),
  FOREIGN KEY (postid)
      REFERENCES posts (id)
);

CREATE TABLE IF NOT EXISTS "categories" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" varchar,
  FOREIGN KEY (id)
      REFERENCES posts (category)
);

CREATE TABLE IF NOT EXISTS "sessions" (
  "sessionid" varchar PRIMARY KEY,
  "userid" int,
  "lastactivity" DATE,
  FOREIGN KEY (userid)
      REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS "message" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "senderid" int,
  "receiverid" int,
  "text" varchar,
  "sentat" DATE,
  "read" int,
  FOREIGN KEY (senderid)
      REFERENCES users (id),
  FOREIGN KEY (receiverid)
      REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS "groups" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "creator" INTEGER,
  "name" text,
  "description" text
);

CREATE TABLE IF NOT EXISTS "usergroups" (
  "uid" INTEGER PRIMARY KEY REFERENCES users(id),
  "gid" INTEGER REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS "events" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "creator" INTEGER,
  "title" text,
  "description" text,
  "time" DATE
);

CREATE TABLE IF NOT EXISTS "userevents" (
  "uid" INTEGER PRIMARY KEY REFERENCES users(id),
  "eid" INTEGER REFERENCES events(id) ON DELETE CASCADE,
  "status" text
);

CREATE TABLE IF NOT EXISTS "groupinvs" (
  "sid" INTEGER REFERENCES users(id),
  "rid" INTEGER REFERENCES users(id),
  "gid" INTEGER REFERENCES groups(id)
);
/* 
UPDATE users
SET online = 0
WHERE username = 123;  */

/*
PRAGMA foreign_keys=off;

 

BEGIN TRANSACTION;

 

ALTER TABLE users RENAME TO _users_old;

 

CREATE TABLE users

(

id INTEGER PRIMARY KEY AUTOINCREMENT,

email varchar UNIQUE NOT NULL,

username varchar UNIQUE NOT NULL,

password varchar NOT NULL,

firstname varchar NOT NULL,

lastname varchar NOT NULL, 

gender varchar NOT NULL,

dateofbirth DATE,

online INTEGER

);

 

INSERT INTO users(id, email, username, password, firstname, lastname, gender, dateofbirth)

SELECT id, email, username, password, firstname, lastname, gender, dateofbirth

FROM _users_old;

 

COMMIT;

 

PRAGMA foreign_keys=on;
*/