/*
Billing System
Textio is making some changes to how they bill users for bulk messages. The system should now calculate the bill based on:

The number of messages sent
The cost per message
Thresholds for discounts
Assignment
Complete the calculateFinalBill and the calculateDiscountRate functions.

calculateFinalBill
Use the calculateBaseBill function to get the cost for the messages sent. Then, use the calculateDiscountRate function to get the discount rate. Finally, calculate the final bill by applying the discount to the base bill and return the result.

The discount is a percentage represented by a float e.g. 10% = .1. Remember that any percentage x% is equal to x / 100.

calculateDiscountRate
This function should take the number of messages sent, and return the relevant discount based on the following discount rates:

10% for more than 1000 messages.
20% for more than 5000 messages.
0% for anything less.
*/

package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?
}

func calculateDiscountRate(messagesSent int) float64 {
	// ?
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
