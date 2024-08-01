create table Users(
  id INT unsigned  NOT NULL AUTO_INCREMENT,
  username VARCHAR(150) NOT NULL,
  password VARCHAR(150) NOT NULL,
  PRIMARY KEY(id)
)

INSERT INTO Users (username, password)
VALUES 
('user1', 'password1'),
('user2', 'password2'),
('user3', 'password3'),
('user4', 'password4'),
('user5', 'password5');

