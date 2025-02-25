CREATE TABLE "task" (
  "id" bigint PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "status" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);