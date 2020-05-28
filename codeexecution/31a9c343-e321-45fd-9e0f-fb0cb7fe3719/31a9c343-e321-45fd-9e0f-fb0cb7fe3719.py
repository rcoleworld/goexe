fizz_buzz = ['fizzbuzz' if i % 3 == 0 and i % 5 == 0 else 'fizz' if i % 3 == 0 else 'buzz' if i % 5 == 0 else i for i in range(100)]
print(fizz_buzz)