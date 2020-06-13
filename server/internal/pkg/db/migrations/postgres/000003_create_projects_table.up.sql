CREATE TABLE IF NOT EXISTS Projects
(
    id          UUID         NOT NULL DEFAULT UUID_GENERATE_V4(),
    title       VARCHAR(127) NOT NULL UNIQUE,
    description VARCHAR(200),
    user_id     UUID,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
)