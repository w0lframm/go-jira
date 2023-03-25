-- Create DB
CREATE DATABASE testdb;
-- Create tables
CREATE TABLE project (
                         id SERIAL PRIMARY KEY,
                         title TEXT
);
CREATE TABLE author (
                        id SERIAL PRIMARY KEY,
                        name TEXT
);
CREATE TABLE statusChange (
                              issueId INTEGER,
                              authorId INTEGER,
                              changeTime TIMESTAMP WITHOUT TIME ZONE,
                              fromStatus TEXT,
                              toStatus TEXT
);
CREATE TABLE issues (
                        id SERIAL PRIMARY KEY,
                        projectId INTEGER,
                        authorId INTEGER,
                        assigneeId INTEGER,
                        key TEXT,
                        summary TEXT,
                        description TEXT,
                        type TEXT,
                        priority TEXT,
                        status TEXT,
                        createdTime TIMESTAMP WITHOUT TIME ZONE,
                        closedTime TIMESTAMP WITHOUT TIME ZONE,
                        updatedTime TIMESTAMP WITHOUT TIME ZONE,
                        timeSpent INTEGER
);

-- Create the pguser user
CREATE USER pguser WITH PASSWORD 'pgpwd';

-- Grant permissions to the pguser user
GRANT USAGE ON SCHEMA public TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE project TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE issues TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE author TO pguser;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE statusChange TO pguser;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO pguser;

-- Create the replicator user and grant replication permission
CREATE USER replicator WITH REPLICATION PASSWORD 'postgres';

-- Grant permissions to the postgres user
GRANT ALL PRIVILEGES ON DATABASE testdb TO postgres;
