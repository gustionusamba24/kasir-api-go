-- Sample data for testing Kasir API
-- Run this SQL script to populate your database with test data

-- Insert sample categories
INSERT INTO categories (name, description) VALUES
('Electronics', 'Electronic devices and accessories'),
('Food & Beverages', 'Food items and drinks'),
('Clothing', 'Apparel and fashion items'),
('Books', 'Books and publications'),
('Home & Garden', 'Home improvement and gardening supplies');

-- Insert sample products
INSERT INTO products (name, price, stock, category_id) VALUES
-- Electronics
('Laptop HP 14"', 7500000.00, 15, 1),
('Wireless Mouse', 150000.00, 50, 1),
('USB-C Cable', 75000.00, 100, 1),
('Bluetooth Speaker', 450000.00, 30, 1),
('Power Bank 10000mAh', 250000.00, 40, 1),

-- Food & Beverages
('Mineral Water 600ml', 3500.00, 200, 2),
('Instant Noodles', 5000.00, 150, 2),
('Coffee Arabica 100g', 45000.00, 50, 2),
('Chocolate Bar', 15000.00, 80, 2),
('Energy Drink', 12000.00, 60, 2),

-- Clothing
('T-Shirt Cotton', 85000.00, 45, 3),
('Jeans Denim', 250000.00, 30, 3),
('Sneakers', 450000.00, 25, 3),
('Cap Baseball', 75000.00, 40, 3),
('Socks Pack of 3', 35000.00, 60, 3),

-- Books
('Programming in Go', 150000.00, 20, 4),
('Database Design', 120000.00, 15, 4),
('Clean Code', 180000.00, 12, 4),
('API Development', 160000.00, 18, 4),

-- Home & Garden
('LED Light Bulb', 25000.00, 100, 5),
('Plant Pot Small', 35000.00, 50, 5),
('Garden Tools Set', 350000.00, 15, 5),
('Cleaning Spray', 28000.00, 70, 5);

-- Insert some products without category (optional/uncategorized)
INSERT INTO products (name, price, stock, category_id) VALUES
('Gift Card Rp 100.000', 100000.00, 200, NULL),
('Promotional Item', 50000.00, 100, NULL);
