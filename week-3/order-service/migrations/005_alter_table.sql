-- ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS status;
-- ALTER TABLE IF EXISTS orders ADD COLUMN IF NOT EXISTS image_url text;
-- ALTER TABLE IF EXISTS orders DROP COLUMN IF EXISTS image_url;
-- ALTER TABLE IF EXISTS products ADD COLUMN IF NOT EXISTS image_url text;
-- ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS role int DEFAULT 0;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT fk_order;
ALTER TABLE IF EXISTS order_items DROP CONSTRAINT fk_product;
