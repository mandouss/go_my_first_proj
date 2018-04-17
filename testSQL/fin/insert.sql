-- insert account
INSERT INTO Account (ACCOUNT_ID, BALANCE)
VALUES ('7942573', '4500.00');

-- insert symbol
INSERT INTO Symbol (SYMBOL_NAME)
VALUES ('SPY');

-- insert accountShare
INSERT INTO AccountShare (ACCOUNT_ID, SYMBOL_ID, SHARE)
VALUES ('7942573', '1', '500.00');

-- insert contractIDtoAccountID
INSERT INTO ContractIDToAccountID (ACCOUNT_ID)
VALUES ('7942573');

-- insert contract
INSERT INTO Contract (CONTRACT_ID, ACCOUNT_ID, SYMBOL_ID, PRICE, AMOUNT, CONTRACT_TYPE)
VALUES ('1', '7942573', '1', 10.11, 15.55, 'part_fill');

-- insert executedContract
INSERT INTO ExecutedContract (CONTRACT_ID, ACCOUNT_ID, SYMBOL_ID, PRICE, AMOUNT)
VALUES ('1', '7942573', '1', 10.11, 15.55);

-- insert CancelledContract
INSERT INTO CancelledContract (CONTRACT_ID, ACCOUNT_ID, SYMBOL_ID, PRICE, AMOUNT)
VALUES ('1', '7942573', '1', 10.11, 15.55);

