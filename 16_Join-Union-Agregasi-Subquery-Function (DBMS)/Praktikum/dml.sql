INSERT INTO transaction_details (id, customer_id, product_id, quantity, price)
VALUES (1, 1, 1, 2, 8000), (2, 2, 2, 1, 10000), (3, 3, 3, 4, 8000);

SELECT customer_name from user where gender = "Lk";

SELECT * from product where product_id = 3;

SELECT * from user where customer_name like "%a%" and created_at >= date_sub(now(), interval 7 day);

SELECT COUNT(user_id) as jumlah from user where gender = "Pr";

SELECT * from user order by customer_name asc;

SELECT * from product limit 5;

UPDATE product SET name_product = "product dummy" WHERE product_id = 1;

UPDATE transaction_details SET quantity = 3 WHERE product_id = 1;

UPDATE product SET deleted_at = NOW() WHERE product_id = 1;

UPDATE product SET deleted_at = NOW() WHERE product_type_id = 1;

DELETE from product where id = 1;

DELETE from product where product_type_id = 1;

SELECT * FROM transaction WHERE id in (1,2);

SELECT sum(price) as price FROM transaction_details WHERE customer_id = 1 group by customer_id;

SELECT customer_id, product_id, sum(price) as price FROM transaction_details group by customer_id, product_id ;

select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;

select * from product p
join product_types pt on pt.id = product_type_id; 

select t.*, p.name_product as nama_product, u.customer_name as user from transaction_details td
join transaction t on t.id = td.transaction_id
join user u on u.id = t.user_id
join product p on p.id = td.product-id
where p.deleted_at is null

delimiter $$
create trigger delete_transaction_detail
before delete on transaction for each row
begin
declare trans_id int;
set trans_id = old.id;
delete from transaction_details where transaction_id = trans_id;
end $$

select * from transaction_details;
delete from transaction where id = 5;

delimiter $$
create trigger update_transaction_id
after delete on transaction_details for each row
begin
declare trans_id int;
set trans_id = old.transaction_id;
update transaction set quanty = (select sum(quanty) from transaction_details where transaction_id = trans_id)
where id = trans_id;
end $$

select * from product where id not in (
select product_id from transaction_details
);
select * from product;






