# OCR Numbers

Welcome to OCR Numbers on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Given a 3 x 4 grid of pipes, underscores, and spaces, determine which number is represented, or whether it is garbled.

## Step One

To begin with, convert a simple binary font to a string containing 0 or 1.

The binary font uses pipes and underscores, four rows high and three columns wide.

```text
     _   #
    | |  # zero.
    |_|  #
         # the fourth row is always blank
```

Is converted to "0"

```text
         #
      |  # one.
      |  #
         # (blank fourth row)
```

Is converted to "1"

If the input is the correct size, but not recognizable, your program should return '?'

If the input is the incorrect size, your program should return an error.

## Step Two

Update your program to recognize multi-character binary strings, replacing garbled numbers with ?

## Step Three

Update your program to recognize all numbers 0 through 9, both individually and as part of a larger string.

```text
 _
 _|
|_

```

Is converted to "2"

```text
      _  _     _  _  _  _  _  _  #
    | _| _||_||_ |_   ||_||_|| | # decimal numbers.
    ||_  _|  | _||_|  ||_| _||_| #
                                 # fourth line is always blank
```

Is converted to "1234567890"

## Step Four

Update your program to handle multiple numbers, one per line.
When converting several lines, join the lines with commas.

```text
    _  _
  | _| _|
  ||_  _|

    _  _
|_||_ |_
  | _||_|

 _  _  _
  ||_||_|
  ||_| _|

```

Is converted to "123,456,789".

## Implementation Notes

Define a function recognizeDigit as described in Step 1 of instructions except make it recognize
all ten digits 0 to 9.  Pick what you like for parameters and return values
but make it useful as a subroutine for step 2.

For Step 2 define,

   func Recognize(string) []string

and implement it using recognizeDigit.

Input strings tested here have a \n at the beginning of each line and
no trailing \n on the last line. (This makes for readable raw string
literals.)

For bonus points, gracefully handle misformatted data.  What should you
do with a partial cell?  Discard it?  Pad with spaces?  Report it with a
"?" character?  What should you do if the first character is not \n?

## Source

### Contributed to by

- @alebaffa
- @bitfield
- @ekingery
- @ferhatelmas
- @hilary
- @kytrinyx
- @leenipper
- @petertseng
- @robphoenix
- @sebito91
- @soniakeys
- @tleen
- @eklatzer

### Based on

Inspired by the Bank OCR kata - https://codingdojo.org/kata/BankOCR/