# sudusolver

Suduko 9x9 puzzle solver using a backtracking method.

## Usage

You can provide the solver with a puzzle matrix encoded in JSON notation like so: -

    `sudusolver -json "[[1,2,3,4,5,6,7,8,9],[4,5,6,7,8,9,1,2,3],[7,8,9,1,2,3,4,5,6],[2,1,4,3,6,5,8,9,7],[3,6,5,8,9,7,2,1,4],[8,9,7,2,1,4,3,6,5],[5,3,1,6,4,2,9,7,8],[6,4,2,0,7,0,5,3,1],[9,7,8,5,0,1,6,4,2]]"`

Unassigned locations contain a 0.

The solution will be returned as JSON: -
    ```json
    [[1,2,3,4,5,6,7,8,9],[4,5,6,7,8,9,1,2,3],[7,8,9,1,2,3,4,5,6],[2,1,4,3,6,5,8,9,7],[3,6,5,8,9,7,2,1,4],[8,9,7,2,1,4,3,6,5],[5,3,1,6,4,2,9,7,8],[6,4,2,9,7,8,5,3,1],[9,7,8,5,3,1,6,4,2]]
    ```