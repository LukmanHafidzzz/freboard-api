CREATE TABLE specifications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT UNIQUE,
    body_material VARCHAR(255),
    neck_material VARCHAR(255),
    fretboard_material VARCHAR(255),
    number_of_frets INT,
    scale_length DOUBLE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);