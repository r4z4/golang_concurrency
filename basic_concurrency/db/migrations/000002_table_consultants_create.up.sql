
CREATE DATABASE  IF NOT EXISTS `basic_golang_concurrency` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `basic_golang_concurrency`;

CREATE TABLE IF NOT EXISTS consultants (
    id SERIAL PRIMARY KEY,
    slug TEXT NOT NULL DEFAULT (uuid_generate_v4()),
    consultant_f_name TEXT NOT NULL,
    consultant_l_name TEXT NOT NULL,
    img_path TEXT DEFAULT NULL,
    -- created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc') NOT NULL,
    -- updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc') NOT NULL
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);