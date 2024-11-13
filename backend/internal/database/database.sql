-- database: :memory:

CREATE TABLE user(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    firstname text   not NULL,
    lastname text not NULL,
    email text not NULL UNIQUE,
    password text not NULL,
    CreateAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    UUID text
);

CREATE TABLE post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    card_id INTEGER,
    FOREIGN KEY (card_id) REFERENCES card(id)
);

CREATE TABLE category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text not NULL
);

INSERT INTO category ("name")
VALUES
    ("Technology"),
    ("Gaming"),
    ("Arts & Entertainment"),
    ("General Discussion");

CREATE TABLE post_category (
    post_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES post(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE card (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    card_id INTEGER,
    post_id INTEGER,
    FOREIGN KEY (card_id) REFERENCES card(id),
    FOREIGN KEY (post_id) REFERENCES post(id)
);

CREATE TABLE likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    card_id INTEGER,
    is_like BOOLEAN,
    FOREIGN KEY (card_id) REFERENCES card(id)
);