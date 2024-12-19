CREATE TABLE Users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  role VARCHAR(20) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE Wallets (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  bank_code VARCHAR(50) NOT NULL,
  bank_account_number VARCHAR(50) NOT NULL,
  bank_account_name VARCHAR(255) not null,
  balance INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);


CREATE TABLE Boys (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  age INT NOT NULL,
  profile_picture_url VARCHAR(200),
  bio TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE Girls (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  age INT NOT NULL,
  profile_picture_url VARCHAR(200),
  bio TEXT,
  daily_rate INT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);


CREATE TABLE Availabilities (
  id SERIAL PRIMARY KEY,
  girl_id INT NOT NULL,
  is_available BOOLEAN NOT NULL,
  start_date DATE,
  end_date DATE,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (girl_id) REFERENCES Girls(id) ON DELETE CASCADE
);

CREATE TABLE Transactions (
  id SERIAL PRIMARY KEY,
  sender_wallet_id INT NOT NULL,
  receiver_wallet_id INT NOT NULL,
  amount INT NOT NULL,
  transaction_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (sender_wallet_id) REFERENCES Wallets(id) ON DELETE CASCADE,
  FOREIGN KEY (receiver_wallet_id) REFERENCES Wallets(id) ON DELETE CASCADE
);

CREATE TABLE Bookings (
  id SERIAL PRIMARY KEY,
  boy_user_id INT NOT NULL,
  girl_user_id INT NOT NULL,
  booking_date DATE NOT NULL,
  num_of_days INT NOT NULL,
  total_cost INT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (boy_user_id) REFERENCES Users(id) ON DELETE CASCADE,
  FOREIGN KEY (girl_user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE Ratings (
  id SERIAL PRIMARY KEY,
  girl_id INT NOT NULL,
  review TEXT,
  stars INT NOT NULL CHECK (stars >= 1 AND stars <= 5),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (girl_id) REFERENCES Girls(id) ON DELETE CASCADE
);

CREATE TABLE User_Activity_logs (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);

CREATE TABLE internal_transactions (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  external_id VARCHAR(255) NOT NULL,
  amount int NOT NULL,
  status VARCHAR(20) not null,
  type VARCHAR(20) not null,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE
);


-- Create 5 users with role 'girls'
INSERT INTO Users (username, password, email, role, created_at, updated_at) 
VALUES 
('lisa_kim', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'lisa.kim@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('mei_chen', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'mei.chen@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('sara_tanaka', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'sara.tanaka@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('maria_santos', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'maria.santos@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('emma_watson', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'emma.watson@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Create their profiles in the Girls table (assuming the users got IDs 1-5)
INSERT INTO Girls (user_id, first_name, last_name, age, profile_picture_url, bio, daily_rate, created_at, updated_at)
VALUES 
(1, 'Lisa', 'Kim', 23, 'https://example.com/lisa.jpg', 'K-pop dance instructor and photography enthusiast. Love traveling and trying new cuisines.', 750000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Mei', 'Chen', 25, 'https://example.com/mei.jpg', 'Professional pianist and tea ceremony expert. Enjoy deep conversations about art and culture.', 850000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Sara', 'Tanaka', 22, 'https://example.com/sara.jpg', 'Aspiring chef and anime lover. Can teach you Japanese cooking and language.', 650000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Maria', 'Santos', 24, 'https://example.com/maria.jpg', 'Salsa dance instructor and beach volleyball player. Always up for adventure!', 800000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Emma', 'Watson', 26, 'https://example.com/emma.jpg', 'Literature graduate and yoga instructor. Love discussing books over coffee.', 900000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Create wallets for each girl
INSERT INTO Wallets (user_id, bank_code, bank_account_number, bank_account_name, balance, created_at, updated_at)
VALUES 
(1, 'BCA', '1234567890', 'Lisa Kim', 1000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'MANDIRI', '2345678901', 'Mei Chen', 1500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'BRI', '3456789012', 'Sara Tanaka', 800000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'BCA', '4567890123', 'Maria Santos', 1200000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'MANDIRI', '5678901234', 'Emma Watson', 2000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Create availability records for all girls
INSERT INTO Availabilities (girl_id, is_available, created_at, updated_at)
VALUES 
(1, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Add some initial ratings
INSERT INTO Ratings (girl_id, review, stars, created_at, updated_at)
VALUES 
(1, 'Amazing dance teacher! Made the experience really fun and engaging.', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(1, 'Great photographer, showed me the best spots in the city.', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'The tea ceremony was a unique and peaceful experience.', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Learned to make perfect sushi! Very patient teacher.', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'Fun personality and great cooking skills.', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'Best salsa instructor ever! So much energy.', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'Incredibly knowledgeable about literature. Great conversations.', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'The yoga session was very relaxing and professional.', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

