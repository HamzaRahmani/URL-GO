CREATE TABLE IF NOT EXISTS url (
    hash char(7) primary key NOT NULL,
		original_url varchar(50) NOT NULL,
		created_at timestamp default current_timestamp
)