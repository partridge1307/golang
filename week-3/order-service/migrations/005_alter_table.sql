ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS status;
ALTER TABLE IF EXISTS orders ADD COLUMN IF NOT EXISTS image_url text;
ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS image_url;
ALTER TABLE IF EXISTS products ADD COLUMN IF NOT EXISTS image_url text;
