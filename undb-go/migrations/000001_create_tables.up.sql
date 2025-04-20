-- 创建表
CREATE TABLE IF NOT EXISTS tables (
    id VARCHAR(36) PRIMARY KEY,
    space_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (space_id) REFERENCES spaces(id) ON DELETE CASCADE
);

-- 创建字段
CREATE TABLE IF NOT EXISTS fields (
    id VARCHAR(36) PRIMARY KEY,
    table_id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    description TEXT,
    required BOOLEAN DEFAULT FALSE,
    unique BOOLEAN DEFAULT FALSE,
    options TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (table_id) REFERENCES tables(id) ON DELETE CASCADE
);

-- 创建记录
CREATE TABLE IF NOT EXISTS records (
    id VARCHAR(36) PRIMARY KEY,
    table_id VARCHAR(36) NOT NULL,
    fields JSON NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (table_id) REFERENCES tables(id) ON DELETE CASCADE
); 