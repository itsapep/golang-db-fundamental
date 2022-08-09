package utils

const (
	INSERT_CUSTOMER = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values (:id,:name,:address,:phone,:email,:saldo)`

	UPDATE_CUSTOMER = `update customer set name=:name, address=:address, phone=:phone, email=:email where id=:id`

	DELETE_CUSTOMER = `update customer set is_status=0 where id=$1`

	INSERT_CUSTOMER_PS = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values ($1,$2,$3,$4,$5,$6)`

	UPDATE_CUSTOMER_PS = `update customer set name=$1, address=$2, phone=$3, email=$4 where id=$5`

	DELETE_CUSTOMER_PS_HD = `delete from customer where id=$1`
	DELETE_CUSTOMER_PS_SD = `update customer set is_status=0 where id=$1`

	SELECT_ALL_CUSTOMER = `select id,name,address,phone,email,saldo
	from customer order by created_at asc limit $1 offset $2`

	SELECT_CUSTOMER_BY_ID = `select id,name,address,phone,email,saldo
	from customer WHERE ID=$1`

	SELECT_CUSTOMER_BY_NAME = `select id,name,address,phone,email,saldo
	from customer WHERE name like $1`

	SELECT_COUNT_CUSTOMER = `select is_status.count(id) as total from customer group by is_status`

	SELECT_COUNT_CUSTOMER_BY_ADDRESS = `select address.count(id) as total from customer group by address`

	SELECT_SUM_SALDO_CUSTOMER = `select sum(saldo) from customer where is_status=1`

	SELECT_ACTIVE_PASSIVE_CUSTOMER = `select select count(id) from customer where is_status=1
										union select count(id) from customer where is_status=0`

	// SHOP
	INSERT_SHOP = `insert into shop 
						(id,no_siup,name,address,phone) 
						values (:id,:no_siup,:name,:address,:phone)`

	UPDATE_SHOP = `update shop set no_siup=:no_siup,name=:name, address=:address, phone=:phone where id=:id`

	DELETE_SHOP = `update shop set is_status=0 where id=$1`

	INSERT_SHOP_PS = `insert into shop
						(id,no_siup,name,address,phone) 
						values ($1,$2,$3,$4,$5)`

	UPDATE_SHOP_PS = `update shop set no_siup=$1, name=$2, address=$3, phone=$4 where id=$5`

	DELETE_SHOP_PS_HD = `delete from shop where id=$1`
	DELETE_SHOP_PS_SD = `update shop set is_status=0 where id=$1`

	SELECT_SHOP_WITH_PRODUCT = `select s.id,s.no_siup,s.name,p.id as product_id,p.name as product_name,p.price,p.stock from shop s join product p on s.id = p.store_id order by no_siup asc limit $1 offset $2`

	// PRODUCT
	INSERT_PRODUCT = `insert into product 
						(id,name,price,description,stock,store_id) 
						values (:id,:name,:price,:description,:stock,:store_id)`

	UPDATE_PRODUCT = `update product set name=:name,price=:price,description=:description,stock=:stock,store_id=:store_id where id=:id`

	DELETE_PRODUCT = `update product set is_status=0 where id=$1`

	INSERT_PRODUCT_PS = `insert into product
						(id,name,price,description,stock,store_id) 
						values ($1,$2,$3,$4,$5,$6)`

	UPDATE_PRODUCT_PS = `update product set name=$1, price=$2, description=$3, stock=$4, store_id=$5 where id=$6`

	DELETE_PRODUCT_PS_HD = `delete from product where id=$1`
	DELETE_PRODUCT_PS_SD = `update product set is_status=0 where id=$1`

	// TRANSACTION
	INSERT_TRANSACTION = `insert into transaction 
						(id, customer_id, product_id, store_id, purchase_date, quantity) 
						values (:id,:no_siup,:name,:address,:phone)`

	UPDATE_TRANSACTION = `update transaction set customer_id=:customer_id, product_id:=product_id, store_id=:store_id, purchase_date=:purchase_date, quantity=:quantity where id=:id`

	DELETE_TRANSACTION = `update transaction set is_status=0 where id=$1`

	INSERT_TRANSACTION_PS = `insert into transaction
	(id, customer_id, product_id, store_id, purchase_date, quantity) 
						values ($1,$2,$3,$4,$5,$6)`

	UPDATE_TRANSACTION_PS = `update transaction set no_customer_id=$1, product_id=$2, store_id=$3, purchase_date=$4, quantity=$5 where id=$6`

	DELETE_TRANSACTION_PS_HD = `delete from transaction where id=$1`
	DELETE_TRANSACTION_PS_SD = `update transaction set is_status=0 where id=$1`
)
