-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name VARCHAR(35),
  second_name VARCHAR(35),
  email VARCHAR(255),
  phone VARCHAR(16),
  password_hash TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS expert_profiles(
  user_id INT REFERENCES users(id) PRIMARY KEY,
  description TEXT,
  preview_photo_url TEXT,
  rating DECIMAL
);

CREATE TABLE IF NOT EXISTS customer_profiles(
  user_id INT REFERENCES users(id) PRIMARY KEY,
  desires TEXT,
  preview_photo_url TEXT,
  rating DECIMAL
);

CREATE TABLE IF NOT EXISTS roles(
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS specialties(
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(128)
);

CREATE TABLE IF NOT EXISTS user_roles(
  user_id INT REFERENCES users(id),
  role_id INT REFERENCES roles(id),

  CONSTRAINT user_role_pkey PRIMARY KEY (user_id, role_id) 
);

CREATE TABLE IF NOT EXISTS expert_specialties(
  expert_id INT REFERENCES expert_profiles(user_id),
  specialty_id INT REFERENCES specialties(id),

  CONSTRAINT expert_specialty_pkey PRIMARY KEY (expert_id, specialty_id) 
);

CREATE TYPE booking_status AS ENUM (
  'pending',
  'accepted',
  'rejected',
  'closed'
);

CREATE TABLE IF NOT EXISTS bookings(
  id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  client_id INT REFERENCES customer_profiles(user_id),
  expert_id INT REFERENCES expert_profiles(user_id),
  price DECIMAL,
  start_time TIMESTAMP,
  end_time TIMESTAMP,
  status booking_status DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS expert_specialties;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS expert_profiles;
DROP TABLE IF EXISTS customer_profiles;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS specialties;
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS booking_status;
-- +goose StatementEnd