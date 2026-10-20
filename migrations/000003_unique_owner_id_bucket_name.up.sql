ALTER TABLE buckets
ADD CONSTRAINT  buckets_owner_name_unique UNIQUE (name, owner_id)