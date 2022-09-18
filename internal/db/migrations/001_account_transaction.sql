CREATE TABLE
    accounts(
        account_id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        type TEXT NOT NULL
    );

CREATE TABLE
    transactions (
        transaction_id INTEGER PRIMARY KEY,
        payee TEXT,
        type TEXT NOT NULL,
        amount REAL DEFAULT 0.0,
        memo TEXT,
        date DATETIME DEFAULT (strftime('%s', 'now')),
        -- date INTEGER DEFAULT (strftime('%s', 'now')),
        account_id INTEGER,
        FOREIGN KEY(account_id) REFERENCES accounts(account_id)
    );