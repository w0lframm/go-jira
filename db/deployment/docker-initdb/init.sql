-- Create tables
CREATE TABLE project (
                         id SERIAL PRIMARY KEY,
                         key TEXT,
                         name TEXT
);
CREATE TABLE author (
                        id SERIAL PRIMARY KEY,
                        key TEXT,
                        name TEXT,
                        display_name TEXT
);
CREATE TABLE issue (
                        id SERIAL PRIMARY KEY,
                        project_id INTEGER,
			            creator_id INTEGER,
                        key TEXT,
                        summary TEXT,
                        type TEXT,
                        priority TEXT,
                        status TEXT
);

-- Grant permissions to the pguser user
GRANT USAGE ON SCHEMA public TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE project TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE issue TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE author TO pguser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO pguser;

-- Create the replicator user and grant replication permission
CREATE USER replicator WITH REPLICATION PASSWORD 'postgres';
