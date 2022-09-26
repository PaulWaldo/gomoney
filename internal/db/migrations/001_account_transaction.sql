CREATE TABLE
    accounts(
        account_id INTEGER PRIMARY KEY,
        name TEXT NOT NULL DEFAULT "",
        type TEXT NOT NULL DEFAULT "",
        memo TEXT NOT NULL DEFAULT "",
        routing TEXT NOT NULL DEFAULT "",
        acct_number TEXT NOT NULL DEFAULT "",
        hidden BOOLEAN DEFAULT FALSE NOT NULL,
        net_worth_include BOOLEAN DEFAULT TRUE NOT NULL,
        budget_include BOOLEAN DEFAULT TRUE NOT NULL
    );

CREATE TABLE
    transactions (
        transaction_id INTEGER PRIMARY KEY,
        payee TEXT NOT NULL,
        type TEXT NOT NULL,
        amount REAL DEFAULT 0.0,
        memo TEXT,
        date DATETIME DEFAULT (strftime('%s', 'now')),
        -- date INTEGER DEFAULT (strftime('%s', 'now')),
        account_id INTEGER,
        FOREIGN KEY(account_id) REFERENCES accounts(account_id)
    );