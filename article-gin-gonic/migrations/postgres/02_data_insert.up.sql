INSERT INTO author (firstname, lastname) VALUES ('Jason', 'Moiron') ON CONFLICT DO NOTHING;
INSERT INTO author (firstname, lastname) VALUES ('John', 'Doe') ON CONFLICT DO NOTHING;
INSERT INTO author (firstname, lastname) VALUES ('Temur', 'Husanov') ON CONFLICT DO NOTHING;

INSERT INTO author (firstname, lastname) VALUES ('Akbar', 'Hasanov') ON CONFLICT DO NOTHING;


INSERT INTO article (title, body, author_id) VALUES ('Lorem1', 'Lorem ipsum1', 1) ON CONFLICT DO NOTHING;
INSERT INTO article (title, body, author_id) VALUES ('Lorem2', 'Lorem ipsum2', 2) ON CONFLICT DO NOTHING;
INSERT INTO article (title, body, author_id) VALUES ('Lorem3', 'Lorem ipsum3', 3) ON CONFLICT DO NOTHING;
INSERT INTO article (title, body, author_id) VALUES ('Lorem4', 'Lorem ipsum4', 4) ON CONFLICT DO NOTHING;
INSERT INTO article (title, body, author_id) VALUES ('Lorem4', 'Lorem ipsum4', 1) ON CONFLICT DO NOTHING;

