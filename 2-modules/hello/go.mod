module example.com/hello

go 1.23.2

replace example.com/greetings => ../greetings

require example.com/bst v0.0.0-00010101000000-000000000000

replace example.com/math => ../math

replace example.com/bst => ../bst
