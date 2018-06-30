# Coin Flips Generator (in Go)

A coin flips generator.

Output format for humans and CSV processors.

That's all folks.

## Usage

```
Usage of ./coinflips:
  -format string
        Output format: 'human', 'csv' (default "human")
  -n uint
        Number of throws (shorthand) (default 10)
  -number uint
        Number of throws (default 10)
  -oneline
        Prints throws on one line
```

### Human Output (multi line)

```
( 1): head     Heads:  1, Tails:  0
( 2): head     Heads:  2, Tails:  0
( 3): head     Heads:  3, Tails:  0
( 4): head     Heads:  4, Tails:  0
( 5): head     Heads:  5, Tails:  0
( 6): tail     Heads:  5, Tails:  1
( 7): head     Heads:  6, Tails:  1
( 8): tail     Heads:  6, Tails:  2
( 9): tail     Heads:  6, Tails:  3
(10): tail     Heads:  6, Tails:  4
```
