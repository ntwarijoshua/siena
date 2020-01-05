CREATE TABLE "public"."profiles"
(
    id SERIAL NOT NULL PRIMARY KEY,
    names VARCHAR(100),
    tag_line TEXT,
    date_of_birth DATE,
    profile_photo TEXT
);