CREATE TABLE IF NOT EXISTS "files" (
                                       "id" SERIAL PRIMARY KEY,
                                       "file_path" varchar(255) NOT NULL,
    "file_name" varchar(255) NOT NULL,
    "file_size_bytes" bigint NOT NULL,
    "file_type" varchar(64) NOT NULL,
    "storage_type" varchar(64) NOT NULL,
    "crc32c" varchar(8) NOT NULL,
    "md5" varchar(32) NOT NULL,
    "metadata" jsonb NOT NULL,
    "created_at" timestamptz DEFAULT now() NOT NULL,
    "updated_at" timestamptz DEFAULT now() NOT NULL,
    "deleted_at" timestamptz
    );
