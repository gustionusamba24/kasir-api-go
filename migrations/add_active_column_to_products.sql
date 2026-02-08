-- Migration: Add active column to products table
-- This migration adds an 'active' column to track product availability status

-- Add active column with default value of true
ALTER TABLE products ADD COLUMN IF NOT EXISTS active BOOLEAN NOT NULL DEFAULT true;

-- Create index on active column for better query performance
CREATE INDEX IF NOT EXISTS idx_products_active ON products(active);

-- Create index on name for case-insensitive search performance
CREATE INDEX IF NOT EXISTS idx_products_name_lower ON products(LOWER(name));

-- Update existing products to be active by default
UPDATE products SET active = true WHERE active IS NULL;
