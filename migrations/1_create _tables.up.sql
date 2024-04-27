CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS coffees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    size DOUBLE PRECISION,
    price_per_unit DOUBLE PRECISION,
    available_units BIGINT,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS extras (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    unit INT NOT NULL,
    price_per_unit DOUBLE PRECISION NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    coffee_id INT NOT NULL,
    user_id INT NOT NULL,
    quantity INT NOT NULL,
    size INT NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (coffee_id) REFERENCES coffees(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS order_extras (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    extra_id INT NOT NULL,
    quantity INT NOT NULL,
    price_per_unit DOUBLE PRECISION NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (extra_id) REFERENCES extras(id)
);
