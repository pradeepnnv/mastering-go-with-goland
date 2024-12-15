CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

select * from todo_items;