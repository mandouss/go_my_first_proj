## 590 hw4 database API design

1. Create an account, return status(suceed or fail)

   * Header: CreateAccount(int ID, double balance) (boolean success)

2. Create a symbol **in an account**,  return status(suceed or fail) 

   * Header: CreateSymbolToAccount(int ID, string symbolName double amount) (boolean success)

2. Buy&sell stock : Create an buy order from an account,  return status. 
   buy: price > 0  sell: price < 0  **NOTE** : amount must be positive

   * Header: CreateOrderToAccount(int ID, string symbolName, double amount, double limit) (int retID, string retSymbolName, double retAmount, double retLimit, boolean success)

4. TBD: Match order, execute order, cancel order

