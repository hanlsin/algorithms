
## Find bigger side of the binary tree
Find the bigger side, "Left" or "Right", of the binary tree.
* Condition1: no calculate a number less than 0
* Condition2: if same sum between left-side and right-side, return empty string.

`ex: {3, 6, 2, 9, -1, 10}`

```text
           3
   --------|-------
   |              |
   6              2
---|---        ---|---
|     |        |
9     -1      10
```

The result is 
```text
Left
```
Because, the sum of left-side is `15`, and the sum of right-side is `12`.

**Solution** is under the "biggersum_of_binarytree" directory.

## Find the largest number in the array
Find the largest number of the array of numbers.
* Condition1: if array is empty, return 0

`ex: {1, 2, 4}`

The results is
```text
4
```

**Solution** is under the "largest_num" directory.

## Find the most profitable occasion
The daily stock prices are the array of numbers like `{6, 0, -1, 10}`.
You can buy when the price is `6` and sell it when the price is `-1`. After then, your profit would be `-7`, it's your loss.

What is the most profitable decision? What is your best profit number?

* Condition1: cannot sell before buying.
* Condition2: if array is empty, return 0

`ex: {6, 0, -1, 10}`

The result is 
```text
11
```
Because, if you buy when the price is `-1` and sell when the price is `10`, `11` would be most profitable.

**Solution** is under the "best_profit" directory.

## Find _k_ th largest number
Which number is the _k_ th largest number in the array?

`ex: {1, 4, 3, 4, 2, 5}`
What is the 3rd largest number?

The result is 
```text
4
```
Because, after the largest number `5`, two `4` numbers are the next largest. So, the 3rd number would be one of this `4` numbers.

**Solution** is under the "kth_number" directory.

## Make a category tree
A category tree is a representation of a set of categories and their parent-child relationships. Each category has a unique name (no two categories have the same name).

1. A category can have a parent category.
2. Categories without a parent category are called root categories.

Example: the category "A" has the categories "B" and "C", then "C" has "D", "E" and "F"...

1. To add a category to a category tree, the name and the parent of the category should be provided. 
2. When adding a root category, a null value should be provided as the parent.
3. A call to getChildren should return the direct children of the specified category in any order.
4. InvalidArgumentException should be thrown when adding a category that has already been added anywhere in the CategoryTree or if a parent is specified but does not exist.

Please write your solution (program to represent an abstract category tree) in pure Golang, JavaScript, PHP, C or C++ and provide examples of using it.

**Solution** is under the "category_tree" directory.