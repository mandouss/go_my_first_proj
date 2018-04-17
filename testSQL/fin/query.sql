-- query in Contract
SELECT Symbol.SYMBOL_NAME, Contract.AMOUNT 
FROM Contract, Symbol
WHERE Contract.SYMBOL_ID=Symbol.SYMBOL_ID AND CONTRACT_ID='1';

-- query in ExecutedContract
SELECT Symbol.SYMBOL_NAME, ExecutedContract.AMOUNT, ExecutedContract.PRICE, EXTRACT(EPOCH FROM ExecutedContract.MODIFICATION_TIME)::integer as TIME_PERIOD
FROM ExecutedContract, Symbol
WHERE ExecutedContract.SYMBOL_ID=Symbol.SYMBOL_ID AND CONTRACT_ID='1';

-- query in CancelledContract
SELECT Symbol.SYMBOL_NAME, CancelledContract.AMOUNT, EXTRACT(EPOCH FROM CancelledContract.MODIFICATION_TIME)::integer as TIME_PERIOD
FROM CancelledContract, Symbol
WHERE CancelledContract.SYMBOL_ID=Symbol.SYMBOL_ID AND CONTRACT_ID='1';


-- prototype I use
-- Select EXTRACT(EPOCH FROM ExecutedContract.MODIFICATION_TIME) from ExecutedContract;
-- SELECT EXTRACT(MINUTE FROM TIMESTAMP '2001-02-16 20:38:40')::integer;