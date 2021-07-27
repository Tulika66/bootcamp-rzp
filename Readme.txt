1. error handling
2. concurrency : mutex & goroutines?
3. retailer has to see his transactions
4. queue : global vs local
          - global makes user dependent on each other
          - local has cost additions

          2 solns
          - hybrid : form cluster of users assigned to one server. since user(of a country):limited
           orders
          - cluster of 200 users having common queue


TODO:
1. Customer should not be able to view quantity of products available
2. Validation of customer , retailer
5. Can induce many-many relation between products and orders ; solves foreign key issue+will implement multiple products
6. Class diagram and flow diagram
7. Orders:Quantity instead of price ?
Fix
8. Instead of validating a product by querying, maintain a set/cache perhaps?
9. Check Clause association and Transaction status
10. Reuse older functions in Order.core while validating productname

// `gorm:"primary_key" sql:"type:ENUM('executed','processing','placed')"`

