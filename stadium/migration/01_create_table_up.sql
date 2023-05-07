
CREATE TYPE location AS ENUM ('long', 'lat', 'how far from user');

CREATE TABLE IF NOT EXISTS stadium (
    id SERIAL PRIMARY KEY,
    stadium_name VARCHAR(100) NOT NULL,
    available_time_start TIMESTAMP,
    available_time_end TIMESTAMP,
    price_p_hour integer not null,
    rating float not null,
    images VARCHAR(50) NOT NULL,
    location jsonb,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


