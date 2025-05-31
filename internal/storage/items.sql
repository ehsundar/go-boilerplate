-- name: CreateItem :one
insert into items (name)
values (@name)
returning id;

-- name: GetItem :one
select *
from items
where id = @id;

-- name: GetItemsPaginated :many
select *
from items
where id < @last_item_id
order by id desc
limit @page_size;
