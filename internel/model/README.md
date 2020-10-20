### model设计原则

1. 数据模型一律嵌入 gorm.Model
2. 数据库主键使用 gorm.Model 提供的ID，业务标识使用Sid
3. 外键关联用ID。如 UserAuth 中的UserId对应User中的ID，而不是Sid
 