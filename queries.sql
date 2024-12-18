-- Users Table
CREATE TABLE Users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  role VARCHAR(20) NOT NULL
);

-- Wallets Table
CREATE TABLE Wallets (
  wallet_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  bank_code VARCHAR(50) NOT NULL,
  bank_account_number VARCHAR(50) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Boys Table
CREATE TABLE Boys (
  boy_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  age INT NOT NULL,
  profile_picture_url VARCHAR(200),
  bio TEXT,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Girls Table
CREATE TABLE Girls (
  girl_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  age INT NOT NULL,
  profile_picture_url VARCHAR(200),
  bio TEXT,
  daily_rate INT NOT null,
  FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Availability Table
CREATE TABLE Availability (
  availability_id SERIAL PRIMARY KEY,
  girl_id INT NOT NULL,
  is_available BOOLEAN NOT NULL,
  FOREIGN KEY (girl_id) REFERENCES Girls(girl_id)
);

-- Transactions Table
CREATE TABLE Transactions (
  transaction_id SERIAL PRIMARY KEY,
  sender_wallet_id INT NOT NULL,
  receiver_wallet_id INT NOT NULL,
  amount INT NOT NULL,
  transaction_date TIMESTAMP NOT NULL,
  FOREIGN KEY (sender_wallet_id) REFERENCES Wallets(wallet_id),
  FOREIGN KEY (receiver_wallet_id) REFERENCES Wallets(wallet_id)
);

-- Bookings Table
CREATE TABLE Bookings (
  booking_id SERIAL PRIMARY KEY,
  boy_id INT NOT NULL,
  girl_id INT NOT NULL,
  booking_date DATE NOT NULL,
  num_of_days INT NOT NULL,
  total_cost INT NOT NULL,
  FOREIGN KEY (boy_id) REFERENCES Boys(boy_id),
  FOREIGN KEY (girl_id) REFERENCES Girls(girl_id)
);

-- Ratings Table
CREATE TABLE Ratings (
  rating_id SERIAL PRIMARY KEY,
  girl_id INT NOT NULL,
  review TEXT,
  stars INT NOT NULL CHECK (stars >= 1 AND stars <= 5),
  FOREIGN KEY (girl_id) REFERENCES Girls(girl_id)
);