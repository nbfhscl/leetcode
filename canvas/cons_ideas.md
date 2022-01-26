# Cons Ideas

## Constraints

### Strings, Arrays and Numbers

1. How many elements can be in the array?
2. what is in each element? If it's a number, is it an interger or a floating
   point?  If it's a string, is it single-byte or multibyte(unicoe)?
3. How large can each element be? If it's a string, how long? If it's a number,
   what is the maximum and maximum value?
4. If the problem involves finding a subsequence, doese "subsequence" mean that
   the elements must be adjacent, or is there no such requirement?
5. Does the array contain unique numbers or can they be repeated (this is
   sometimes relevent)?

### Return Values

1. What should my method return? For example, if I’m trying to find the longest
   subsequence of increasing numbers in an array, should I return the length,
   the start index, or both?
2. If there are multiple solutions to the problem, which one should be returned?
3. If it should return multiple values, do you have any preference on what to
   return? E.g. should it return an object, a tuple, an array, or pass the
   output parameters as input references? (This may not be applicable in
   languages allowing you to return multiple values, e.g. Python)
4. What should I do/return if the input is invalid / does not match the
   constraints? Options may be to do nothing (always assume the input is
   correct), raise an exception, or return some specific value.
5. In problems where you’re supposed to find something (e.g. a number in an
   array), what should be returned if the element is not present?

### Grids/Mazes

1. For problems where some actor (e.g. a robot) is moving in a grid or maze,
   what moves are allowd? Can the robot move diagonally (hence 8 valid moves),
   or only horizontally/vertically (hence only 4 valid moves)?
2. Are all cells in the grid allowed, or can there be obstacles?
3. If the actor is trying to get from cell A to cell B, are cells A and B
   guaranteed to be different from each other?
4. If the actor is trying to get from cell A to cell B, is it guaranteed that
   there’s a path between the two cells?

### Graphs

1. How many nodes can the graph have?
2. How many edges can the graph have?
3. If the edges have weights, what is the range for the weights?
4. Can there be loops in the graph? Can there be negative-sum loops in the
   graph?
5. Is the graph directed or undirected?
6. Does the graph have multiple edges and/or self-loops?

## Ideas

1. complexities. time:, memory:.
2. complexities. time:, memory:.

