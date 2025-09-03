CREATE TYPE debt_transaction_type AS ENUM ('LEND', 'REPAY');
CREATE TYPE debt_transaction_action AS ENUM ('LEND', 'BORROW', 'RECEIVE', 'RETURN');

CREATE TABLE IF NOT EXISTS transfer_methods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    display TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS debt_transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lender_profile_id UUID NOT NULL,
    borrower_profile_id UUID NOT NULL,
    type debt_transaction_type NOT NULL,
    action debt_transaction_action NOT NULL,
    amount NUMERIC(20, 2) NOT NULL,
    transfer_method_id UUID NOT NULL REFERENCES transfer_methods(id),
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS debt_transactions_lender_profile_id_idx ON debt_transactions(lender_profile_id);
CREATE INDEX IF NOT EXISTS debt_transactions_borrower_profile_id_idx ON debt_transactions(borrower_profile_id);
CREATE INDEX IF NOT EXISTS debt_transactions_transfer_method_id_idx ON debt_transactions(transfer_method_id);
CREATE INDEX IF NOT EXISTS debt_transactions_created_at_idx ON debt_transactions(created_at);
