-- Create tables
CREATE TABLE project (
                         id SERIAL PRIMARY KEY,
                         key TEXT,
                         name TEXT,
                         url TEXT
);
CREATE TABLE author (
                        id SERIAL PRIMARY KEY,
                        key TEXT,
                        name TEXT,
                        display_name TEXT
);
CREATE TABLE issue (
                       id SERIAL PRIMARY KEY,
                       project_id INTEGER REFERENCES project (id) ON DELETE CASCADE,
                       creator_id INTEGER REFERENCES author (id) ON DELETE CASCADE,
                       assignee_id INTEGER REFERENCES author (id) ON DELETE CASCADE,
                       key TEXT,
                       summary TEXT,
                       type TEXT,
                       priority TEXT,
                       status TEXT,
                       created TIMESTAMP WITHOUT TIME ZONE,
                       updated TIMESTAMP WITHOUT TIME ZONE,
                       timespent INTEGER
);
CREATE TABLE status_change (
                              issue_id INTEGER REFERENCES issue (id) ON DELETE CASCADE,
                              author_id INTEGER REFERENCES author (id) ON DELETE CASCADE,
                              changed TIMESTAMP WITHOUT TIME ZONE,
                              from_status TEXT,
                              to_status TEXT
);

-- Grant permissions to the pguser user
GRANT USAGE ON SCHEMA public TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE project TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE issue TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE author TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE status_change TO pguser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO pguser;

-- Create the replicator user and grant replication permission
CREATE USER replicator WITH REPLICATION PASSWORD 'postgres';

-- Grant permissions to the postgres user
GRANT ALL PRIVILEGES ON DATABASE testdb TO postgres;
