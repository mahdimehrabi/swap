CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP
);

CREATE TABLE coins (
                       id SERIAL PRIMARY KEY,
                       symbol VARCHAR(20) NOT NULL UNIQUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP

);

CREATE TABLE coin_users (
                            coin_id INTEGER NOT NULL,
                            user_id INTEGER NOT NULL,
                            amount VARCHAR(255) NOT NULL ,
                            PRIMARY KEY (coin_id, user_id),
                            CONSTRAINT fk_coin
                                FOREIGN KEY(coin_id)
                                    REFERENCES coins(id)
                                    ON DELETE CASCADE,
                            CONSTRAINT fk_user
                                FOREIGN KEY(user_id)
                                    REFERENCES users(id)
                                    ON DELETE CASCADE
);

CREATE INDEX idx_coin_users_coin_id ON coin_users (coin_id);
CREATE INDEX idx_coin_users_user_id ON coin_users (user_id);