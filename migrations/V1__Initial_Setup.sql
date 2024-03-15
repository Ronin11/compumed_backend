CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE devices(
   id             uuid DEFAULT uuid_generate_v4 (),
   user_id        uuid,
   name           varchar,
--    created_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--    data           jsonb,
   PRIMARY KEY (id)
)
