CREATE TABLE IF NOT EXISTS movie (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(200) NOT NULL,
    year date not null,
    actor_id integer,
    rating float not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    foreign key (actor_id) references actor(id)
);

CREATE TABLE IF NOT EXISTS actor (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    age int NOT NULL,
    country VARCHAR(45),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);





-- create table if not exists movies_actors(
--     id SERIAL PRIMARY KEY,
--     movie_id integer NOT NULL,
--     actor_id integer NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     foreign key (movie_id) references movie(id),
--     foreign key (actor_id) references actor(id)
-- );
