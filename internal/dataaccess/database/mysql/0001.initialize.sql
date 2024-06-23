CREATE TABLE IF NOT EXISTS users (
    id UNSIGNED BIGINT AUTO_INCREMENT,
    user_name VARCHAR(256),
    PRIMARY KEY (id),
    UNIQUE (user_name)
);

CREATE TABLE IF NOT EXISTS accounts (
    of_user_id UNSIGNED BIGINT,
    hash VARCHAR(128) NOT NULL,
    PRIMARY KEY (of_user_id),
    FOREIGN KEY (of_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS polls (
    id UNSIGNED BIGINT AUTO_INCREMENT,
    of_user_id UNSIGNED BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (of_user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS questions (
    id UNSIGNED BIGINT AUTO_INCREMENT,
    of_poll_id UNSIGNED BIGINT,
    question VARCHAR(2048),
    answer VARCHAR(2048),
    PRIMARY KEY (id),
    FOREIGN KEY (of_poll_id) REFERENCES polls(id)
);
