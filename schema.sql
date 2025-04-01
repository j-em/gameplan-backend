CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    stytchId varchar(40) NOT NULL,
    stripeId varchar(30) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(20),
    country varchar(2),
    birthday INTEGER,
    lang VARCHAR(2) NOT NULL DEFAULT 'en' CHECK (lang IN ('fr', 'en')),
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    isActive boolean NOT NULL DEFAULT true,
    isVerified boolean NOT NULL DEFAULT false,
    subscriptionTier VARCHAR(4) NOT NULL DEFAULT 'free' CHECK (subscriptionTier IN ('free', 'pro')),
    jsonSettings TEXT,
    UNIQUE (email)
);

CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    userId INTEGER REFERENCES users (id),
    name varchar(255) NOT NULL,
    email varchar(255),
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    preferredMatchGroup integer,
    isActive boolean NOT NULL DEFAULT true,
    emailNotificationsEnabled boolean NOT NULL DEFAULT false,
    UNIQUE (name)
);

CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    userId INTEGER REFERENCES users (id),
    name varchar(255) NOT NULL,
    startDate date NOT NULL,
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    isActive boolean NOT NULL DEFAULT true,
    seasonType varchar(50) CHECK (
        seasonType IN (
            'pool',
            'bowling',
            'other'
        )
    ) NOT NULL,
    frequency varchar(50) CHECK (
        frequency IN (
            'weekly',
            'biweekly',
            'monthly',
            'quarterly',
            'yearly'
        )
    ) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    seasonId integer REFERENCES seasons (id),
    playerId1 integer REFERENCES players (id),
    playerId1Points integer NOT NULL,
    playerId2 integer REFERENCES players (id),
    playerId2Points integer NOT NULL,
    matchDate date NOT NULL,
    winnerId integer REFERENCES players (id),
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    isActive boolean NOT NULL DEFAULT true,
    "group" integer NOT NULL
);

CREATE TABLE player_custom_columns (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    field_type VARCHAR(50) NOT NULL,
    description TEXT,
    is_required BOOLEAN DEFAULT false,
    is_active BOOLEAN DEFAULT true,
    display_order INTEGER,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name)
);

CREATE TABLE player_custom_values (
    id SERIAL PRIMARY KEY,
    player_id INTEGER REFERENCES players (id),
    column_id INTEGER REFERENCES player_custom_columns (id),
    value TEXT,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (player_id, column_id)
);

CREATE TABLE match_custom_columns (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    field_type VARCHAR(50) NOT NULL,
    description TEXT,
    is_required BOOLEAN DEFAULT false,
    is_active BOOLEAN DEFAULT true,
    display_order INTEGER,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (name)
);

CREATE TABLE match_custom_values (
    id SERIAL PRIMARY KEY,
    match_id INTEGER REFERENCES matches (id),
    column_id INTEGER REFERENCES match_custom_columns (id),
    value TEXT,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (match_id, column_id)
);
