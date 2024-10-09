--customers
INSERT INTO customer (customer_id, name, phone, address) VALUES ([customer_id], [name], [phone], [address]);
UPDATE customer SET name=[name], phone=[phone], address=[address], updated_at=NOW() WHERE customer_id=[customer_id];
DELETE FROM customer WHERE customer_id=[customer_id];



-- services
INSERT INTO service (service_id, service_name, unit, price) VALUES ([service_id],[service_name], [unit], [price]);
UPDATE service SET service_name=[service_name], unit=[unit], price=[price], updated_at=NOW() WHERE service_id=[service_id];
DELETE FROM service WHERE service_id=[service_id];

--order
INSERT INTO "order" (order_id, customer_id, order_date, received_by) VALUES ([order_id], [customer_id], NOW(), [received_by]);
UPDATE "order" SET completion_date=[completion_date], updated_at=NOW() WHERE order_id=[order_id];

-- order detail
INSERT INTO order_detail (order_id, service_id, qty)VALUES ([order_id], [service_id], [qty]);
