CREATE TABLE "Users" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "Urls" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid,
  "shortened_url" varchar UNIQUE NOT NULL,
  "original_url" varchar NOT NULL,
  "click_count" bigint NOT NULL DEFAULT (0),
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "Users" ("username");

CREATE INDEX ON "Urls" ("shortened_url");

CREATE INDEX ON "Urls" ("user_id");

ALTER TABLE "Urls" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");