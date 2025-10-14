

Base URL: `http://localhost:8080/`

---
# 接口名 `[请求类型] [接口]`

**描述**：无

**请求体**：无

**返回示例**：无

---

# 获取区块链 `GET /api/v1/get_blocks`

**描述**：获取整个区块链及当前交易池信息

**请求体**：无

**返回示例**：

```json
{
  "chain": [
    {
      "index": 0,
      "timestamp": 1697080000,
      "transactions": [],
      "prev_hash": "0",
      "hash": "abcdef123456...",
      "nonce": 0,
      "difficulty": 3
    },
    {
      "index": 1,
      "timestamp": 1697080100,
      "transactions": [
        { "from": "Alice", "to": "Bob", "amount": 10 }
      ],
      "prev_hash": "abcdef123456...",
      "hash": "123456abcdef...",
      "nonce": 0,
      "difficulty": 3
    }
  ],
  "state": null,
  "difficulty": 3,
  "pending_tx": [
    { "from": "Charlie", "to": "Dave", "amount": 5 }
  ]
}
```

---

# 打包交易 `POST /api/v1/post_blocks`

**描述**：根据交易池打包新区块并添加到区块链

**请求体**：

```json
{
  "pending_tx": [
    { "from": "Alice", "to": "Bob", "amount": 10 },
    { "from": "Charlie", "to": "Dave", "amount": 5 }
  ]
}
```

**返回示例**：

```json
{
  "index": 2,
  "timestamp": 1697080200,
  "transactions": [
    { "from": "Alice", "to": "Bob", "amount": 10 },
    { "from": "Charlie", "to": "Dave", "amount": 5 }
  ],
  "prev_hash": "123456abcdef...",
  "hash": "7890abcdef123...",
  "nonce": 0,
  "difficulty": 3
}
```

---

# 获取交易池 `GET /api/v1/get_transactions`

**描述**：查看当前待打包的交易

**请求体**：无

**返回示例**：

```json
[
  { "from": "Alice", "to": "Bob", "amount": 10 },
  { "from": "Charlie", "to": "Dave", "amount": 5 }
]
```

---

# 挖矿生成新区块 `POST /api/v1/mine`

**描述**：将交易池中的交易打包生成新区块，清空交易池

**请求体**：无

**返回示例**：

```json
{
  "index": 4,
  "timestamp": 1697080400,
  "transactions": [
    { "from": "Alice", "to": "Bob", "amount": 10 },
    { "from": "Charlie", "to": "Dave", "amount": 5 }
  ],
  "prev_hash": "abc123def456...",
  "hash": "def456abc123...",
  "nonce": 0,
  "difficulty": 3
}
```

**错误示例**（无交易可打包）：

```json
{
  "error": "No transactions to mine"
}
```

---

# 获取所有账户余额 `[GET] /api/v1/accounts`

**描述**：获取系统中所有账户及其当前余额。

**请求体**：无

**返回示例**：

```json
[
  {
    "address": "alice",
    "balance": 120.5
  },
  {
    "address": "bob",
    "balance": 80.0
  },
  {
    "address": "miner01",
    "balance": 50.0
  }
]
```

---

# 获取单个账户余额 `[GET] /api/v1/accounts/:address`

**描述**：根据地址查询指定账户的余额。

**请求体**：无

**返回示例**：

```json
{
  "address": "alice",
  "balance": 120.5
}
```

---

# 提交交易 `[POST] /api/v1/accounts/transaction`

**描述**：从一个账户向另一个账户转账（简化的 UTXO 模型）。

**请求体**：

```json
{
  "from": "alice",
  "to": "bob",
  "amount": 25.0
}
```

**返回示例**：

```json
{
  "message": "Transaction applied",
  "transaction": {
    "txid": "5f3a9fcd-8d43-4ac2-9404-93b5aeb3a1b7",
    "from": "alice",
    "to": "bob",
    "amount": 25
  }
}
```

**错误示例**（余额不足）：

```json
{
  "error": "Insufficient balance"
}
```

---
