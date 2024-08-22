-- Create "passengers" table
CREATE TABLE "public"."passengers" (
  "id" bigint NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "phone_number" character varying(16) NULL,
  "vne_id" character varying(16) NULL,
  "name" text NULL,
  "dob" timestamptz NULL,
  "gender" text NULL,
  "personal_image" text NULL,
  "account_type" text NULL,
  "confirmation_document" text NULL,
  "status" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_passengers_deleted_at" to table: "passengers"
CREATE INDEX "idx_passengers_deleted_at" ON "public"."passengers" ("deleted_at");
-- Create index "idx_unique_passenger" to table: "passengers"
CREATE UNIQUE INDEX "idx_unique_passenger" ON "public"."passengers" ("phone_number", "vne_id");
