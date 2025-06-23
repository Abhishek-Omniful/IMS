

-- Create Product Table
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    general_description TEXT,
    seller_id BIGINT NOT NULL REFERENCES sellers(id) ON DELETE CASCADE
);

-- Create Tenant Table
CREATE TABLE tenants (
    id BIGSERIAL PRIMARY KEY,
    tenant_name TEXT NOT NULL,
    registered_address TEXT NOT NULL,
    tenant_contact TEXT NOT NULL,
    tenant_email TEXT NOT NULL UNIQUE
);

-- Create Hub Table
CREATE TABLE hubs (
    id BIGSERIAL PRIMARY KEY,
    --If a tenant is deleted from the tenants table, then all hubs rows linked to that tenant are automatically deleted.
    tenant_id BIGINT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    manager_name TEXT NOT NULL,
    manager_contact TEXT NOT NULL,
    manager_email TEXT NOT NULL
);

-- Create Seller Table
CREATE TABLE sellers (
    id BIGSERIAL PRIMARY KEY,
    hub_id BIGINT NOT NULL REFERENCES hubs(id) ON DELETE CASCADE,
    tenant_id BIGINT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    seller_name TEXT NOT NULL,
    seller_contact TEXT NOT NULL,
    seller_email TEXT NOT NULL
);

-- Create SKU Table
CREATE TABLE skus (
    id BIGSERIAL PRIMARY KEY,
    seller_id BIGINT NOT NULL REFERENCES sellers(id) ON DELETE CASCADE,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    images TEXT,
    description TEXT,
    fragile BOOLEAN DEFAULT FALSE,
    dimensions TEXT
);

-- Create Address Table
CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY,
    entity_id BIGINT NOT NULL,
    entity_type TEXT NOT NULL CHECK (entity_type IN ('tenant', 'hub', 'seller')),
    address_line1 TEXT NOT NULL,
    address_line2 TEXT,
    pincode TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL
);


-- Table: validate_order_requests
CREATE TABLE IF NOT EXISTS validate_order_requests (
    id SERIAL PRIMARY KEY,
    sku_id VARCHAR NOT NULL,
    hub_id VARCHAR NOT NULL
);

-- Table: validation_responses
CREATE TABLE IF NOT EXISTS validation_responses (
    id SERIAL PRIMARY KEY,
    is_valid BOOLEAN NOT NULL,
    error TEXT
);

--  inventories


CREATE TABLE inventories (
    sku_id BIGINT NOT NULL,
    hub_id BIGINT NOT NULL,
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    unit_price INT NOT NULL DEFAULT 0 CHECK (unit_price >= 0),
    PRIMARY KEY (sku_id, hub_id),
    FOREIGN KEY (sku_id) REFERENCES skus(id),
    FOREIGN KEY (hub_id) REFERENCES hubs(id)
);

