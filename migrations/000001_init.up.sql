CREATE TABLE word_list (
    ID SERIAL PRIMARY KEY,
    word TEXT NOT NULL,
    is_filtred BOOL
);

CREATE TABLE url_list (
    ID SERIAL PRIMARY KEY,
    link TEXT NOT NULL,
    is_parsed BOOL
);

CREATE TABLE word_location (
    ID SERIAL PRIMARY KEY,
    fk_word_ID INT REFERENCES word_list(ID),
    fk_url_ID INT REFERENCES url_list(ID),
    location INT
);

CREATE TABLE link_between_url (
    ID SERIAL PRIMARY KEY,
    fk_from_url_ID INT REFERENCES url_list(ID),
    fk_to_url_ID INT REFERENCES url_list(ID)
);

CREATE TABLE link_word (
    ID SERIAL PRIMARY KEY,
    fk_wordId INT REFERENCES word_list(ID),
    fk_linkId INT REFERENCES link_between_url(ID)
);
