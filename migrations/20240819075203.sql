-- Create "accounts" table
CREATE TABLE "public"."accounts" (
  "id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "username" text NULL,
  "password" text NULL,
  "name" text NULL,
  "role" text NULL,
  "personal_image" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX "idx_accounts_deleted_at" ON "public"."accounts" ("deleted_at");
-- Create "refresh_tokens" table
CREATE TABLE "public"."refresh_tokens" (
  "account_id" bigint NOT NULL,
  "token" text NOT NULL,
  "expiration" timestamptz NULL,
  PRIMARY KEY ("account_id", "token"),
  CONSTRAINT "fk_accounts_refresh_tokens" FOREIGN KEY ("account_id") REFERENCES "public"."accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_refresh_tokens_token" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_token" ON "public"."refresh_tokens" ("token");
