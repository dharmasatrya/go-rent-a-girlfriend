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

CREATE TABLE Availability (
  id SERIAL PRIMARY KEY,
  girl_id INT NOT NULL,
  is_available BOOLEAN NOT NULL,
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
  boy_id INT NOT NULL,
  girl_id INT NOT NULL,
  booking_date DATE NOT NULL,
  num_of_days INT NOT NULL,
  total_cost INT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (boy_id) REFERENCES Boys(id) ON DELETE CASCADE,
  FOREIGN KEY (girl_id) REFERENCES Girls(id) ON DELETE CASCADE
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

-- Insert Users
INSERT INTO Users (username, password, email, role, created_at, updated_at)
VALUES 
('aisha123', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'aisha@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('sarah123', '$2a$10$xYu0qTMR7qR7g/z1Z8m5huqX7ZyB2.PqHBk.QiAXwBvTQ9kh6vtmO', 'sarah@example.com', 'girls', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert Girls (using the user IDs from above)
INSERT INTO Girls (user_id, first_name, last_name, age, profile_picture_url, bio, daily_rate, created_at, updated_at)
VALUES 
(1, 'Aisha', 'Singgih', 23, 'https://example.com/aisha.jpg', 'Love to travel and meet new people', 500000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Sarah', 'Putri', 25, 'https://example.com/sarah.jpg', 'Enjoy cooking and having deep conversations', 450000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert Wallets (using the user IDs from above)
INSERT INTO Wallets (user_id, bank_code, bank_account_number, balance, created_at, updated_at)
VALUES 
(1, 'BCA', '1234567890', 1000000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'MANDIRI', '0987654321', 800000, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);