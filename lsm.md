# LSM
LSM 树是用牺牲读取性能来尽可能换取写入性能和空间利用率.

LSM 高速写入的特性对分布式数据库而言是有非常大吸引力的，同时其KV 结构更是分片所喜欢的一种数据格式，非常适合基于此构建分布式数据库. 所以诸如 Apache Cassandra、ClickHouse、OceanBase、TiDB 等分布式数据库都选用 LSM 树或类似结构的存储引擎来构建分布式数据库.

## compaction
常见的合并策略有 Size-Tiered Compaction 和 Leveled Compaction.

经典的 RocksDB 会在 L0 合并到 L1 时，使用 Size-Tiered Compaction；而从 L1 开始，则是采用经典的 Leveled Compaction, 这其中原因是 L0 的数据表之间肯定会存在相同的 key.

## FAQ
### 放大
- 读放大: 来源于在读取时需要在多个文件中获取数据并解决数据冲突问题，如查询操作中所示的，读取的目标越多，对读取操作的影响越大，而合并操作可以有效缓解读放大问题。
- 写放大: 对于 LSM 树来说，写放大来源于持续的合并操作，特别是 Leveled Compaction，可以造成多层连续进行合并操作, 这样会让写放大问题呈几何倍增长
- 空间放大: 相同 key 的数据被放置了多份，这是在合并操作中所产生的, 尤其是 Size-Tiered Compaction 会有严重的空间放大问题