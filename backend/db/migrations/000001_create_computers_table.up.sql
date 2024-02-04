CREATE TABLE computers (
  "id" SERIAL NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMPTZ,
  "host_name" varchar NOT NULL,
  "ip_address" varchar(15) NOT NULL,
  "mac_address" varchar(17) NOT NULL
)