-- Insert some initial customers
INSERT INTO customer (name, phone, address) VALUES
('Alice', '123456789', '123 Wonderland Ave'),
('Bob', '987654321', '456 Nowhere Blvd');

-- Insert some initial services
INSERT INTO service (service_name, unit, price) VALUES
('Wash and Fold', 'Kg', 5000),
('Dry Cleaning', 'Piece', 20000),
('Ironing', 'Piece', 5000);

-- Insert sample orders
INSERT INTO "order" (customer_id, order_date, received_by) VALUES
(1, '2024-08-15', 'John'),
(2, '2024-08-16', 'Jane');

-- Insert order details for the above orders
INSERT INTO order_detail (order_id, service_id, qty) VALUES
(1, 1, 10),
(1, 2, 2),
(2, 1, 5);
