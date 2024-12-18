## defer

in go lang code execution happens line by line. But when go complier finds defer it executes the defer after all the code in surrounding is executed.

basically it follows lifo style to execute defer code.