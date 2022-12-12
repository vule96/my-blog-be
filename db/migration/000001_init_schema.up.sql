CREATE TABLE "posts" (
  "id" varchar primary key not null,
  "user_id" varchar not null,
  "content" varchar not null,
  "is_active" boolean not null default true,
  "created_at" timestamptz not null default (now()),
  "updated_at" timestamptz not null default (now())
);

CREATE INDEX ON "posts" ("user_id");