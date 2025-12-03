CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  role VARCHAR(20) DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS idc_info (
  id SERIAL PRIMARY KEY,
  zbx_id VARCHAR(50) UNIQUE NOT NULL,
  idc_code VARCHAR(10) NOT NULL,
  idc_name VARCHAR(50) NOT NULL,
  ipmi_ip VARCHAR(16) NOT NULL,
  ssh_ip VARCHAR(16) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS machine_info (
  id SERIAL PRIMARY KEY,
  zbx_id VARCHAR(50) UNIQUE NOT NULL,
  system_type VARCHAR(20) NOT NULL,
  manufacturer VARCHAR(20) NOT NULL,
  server_sn VARCHAR(50) NOT NULL,
  system_disk VARCHAR(20) NOT NULL,
  ssd_count VARCHAR(20) NOT NULL,
  hdd_count VARCHAR(20) NOT NULL,
  memory_count VARCHAR(20) NOT NULL,
  cpu_info TEXT NOT NULL,
  server_height VARCHAR(10) NOT NULL,
  CONSTRAINT fk_machine_info_zbx_id FOREIGN KEY (zbx_id) REFERENCES idc_info(zbx_id)
);


CREATE TABLE IF NOT EXISTS business_info (
  id SERIAL PRIMARY KEY,
  zbx_id VARCHAR(50) UNIQUE NOT NULL,
  business_name VARCHAR(100) NOT NULL,
  business_id VARCHAR(50) NOT NULL,
  old_business_name VARCHAR(100) NOT NULL,
  old_business_id VARCHAR(50) NOT NULL,
  business_speed SMALLINT NOT NULL,
  old_business_speed SMALLINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_business_info_zbx_id FOREIGN KEY (zbx_id) REFERENCES idc_info(zbx_id)
);


CREATE TABLE IF NOT EXISTS network_info (
  id SERIAL PRIMARY KEY,
  zbx_id VARCHAR(50) UNIQUE NOT NULL,
  mac_address VARCHAR(17) NOT NULL,
  eth_name VARCHAR(15) NOT NULL,
  idc_code VARCHAR(10) NOT NULL,
  net_type VARCHAR(20) NOT NULL,
  vlan VARCHAR(9) NOT NULL,
  ipv4_ip VARCHAR(20) NOT NULL,
  ipv4_gataway VARCHAR(20) NOT NULL,
  ipv6_ip VARCHAR(50),
  ipv6_gataway VARCHAR(50),
  ip_speed SMALLINT NOT NULL,
  ip_status VARCHAR(10),
  ip_notes VARCHAR(255),
  segment_notes VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS version_info (
  version_num VARCHAR(20) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_machine_info_zbx_id ON machine_info(zbx_id);
CREATE INDEX IF NOT EXISTS idx_business_info_zbx_id ON business_info(zbx_id);
CREATE INDEX IF NOT EXISTS idx_idc_info_zbx_id ON idc_info(zbx_id);
CREATE INDEX IF NOT EXISTS idx_network_info_zbx_id ON network_info(zbx_id);


INSERT INTO users (username, password_hash, role) VALUES 
('admin', '$2a$10$4bN3oqY8z8z8z8z8z8z8zO1234567890123456789012345678901234', 'admin'),
('bdkj', '$2a$10$4bN3oqY8z8z8z8z8z8z8zO1234567890123456789012345678901234', 'user')
ON CONFLICT (username) DO NOTHING;
