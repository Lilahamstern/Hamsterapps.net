CREATE TABLE IF NOT EXISTS Projects
(
    id          BIGSERIAL    NOT NULL,
    title       VARCHAR(127) NOT NULL UNIQUE,
    description VARCHAR(200),
    user_id     BIGSERIAL,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
)