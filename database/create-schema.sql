CREATE TABLE item
(
    id        SERIAL PRIMARY KEY,
    collector TEXT NOT NULL,
    name      TEXT NOT NULL
);

CREATE TABLE tag
(
    id        SERIAL PRIMARY KEY,
    collector TEXT NOT NULL,
    name      TEXT NOT NULL
);

CREATE TABLE item_tag
(
    item_id INTEGER REFERENCES item,
    tag_id  INTEGER REFERENCES tag,
    PRIMARY KEY (item_id, tag_id)
);