# Data Structures and Algorithms Curriculum (Implementation Focused)

This curriculum is designed to help you learn Data Structures and Algorithms by implementing them yourself.

## Foundational Data Structures

1.  **Arrays:**
    * Basic array implementation (fixed size)
    * Dynamic array/resizable array implementation
    * Implement common array operations: insertion, deletion, searching (linear), updating.
    * **Resources:**
        * [GeeksforGeeks - Array](https://www.geeksforgeeks.org/array-data-structure/)
        * [Khan Academy - Arrays](https://www.khanacademy.org/computing/computer-programming/programming-natural-simulations/programming-arrays/v/what-are-arrays)
        * [Coursera - Data Structures and Algorithms Specialization (University of California, San Diego)](https://www.coursera.org/specializations/data-structures-algorithms) (General resource, specific modules on arrays)

2.  **Linked Lists:**
    * Singly Linked List: Implement insertion (at head, tail, specific index), deletion (at head, tail, specific index), searching, reversing.
    * Doubly Linked List: Implement the same operations as above, along with backward traversal.
    * Circular Linked List: Understand its properties and implement basic operations.
    * **Resources:**
        * [GeeksforGeeks - Linked List](https://www.geeksforgeeks.org/data-structures/linked-list/)
        * [Khan Academy - Linked Lists](https://www.khanacademy.org/computing/computer-science/algorithms/linked-lists/v/introduction-to-linked-lists)
        * [MIT OpenCourseware - Introduction to Algorithms (SMA 05503)](https://ocw.mit.edu/courses/6-006-introduction-to-algorithms-sma-05503-fall-2005/pages/lecture-notes/) (Lecture notes often cover linked lists)

3.  **Stacks:**
    * Implement using an array.
    * Implement using a linked list.
    * Implement the `push`, `pop`, `peek`, and `isEmpty` operations.
    * **Resources:**
        * [GeeksforGeeks - Stack](https://www.geeksforgeeks.org/stack-data-structure/)
        * [Khan Academy - Stacks](https://www.khanacademy.org/computing/computer-science/algorithms/stacks-and-queues/v/stacks-and-queues)

4.  **Queues:**
    * Implement using an array (linear queue and circular queue).
    * Implement using a linked list.
    * Implement the `enqueue`, `dequeue`, `peek`, and `isEmpty` operations.
    * **Resources:**
        * [GeeksforGeeks - Queue](https://www.geeksforgeeks.org/queue-data-structure/)
        * [Khan Academy - Queues](https://www.khanacademy.org/computing/computer-science/algorithms/stacks-and-queues/v/stacks-and-queues)

5.  **Hash Tables (or Hash Maps):**
    * Implement using an array and a collision resolution strategy (e.g., separate chaining with linked lists or open addressing with linear probing).
    * Implement `insert`, `delete`, `search` operations.
    * **Resources:**
        * [GeeksforGeeks - Hash Table](https://www.geeksforgeeks.org/hashing-data-structure/)
        * [Khan Academy - Hash Tables](https://www.khanacademy.org/computing/computer-science/data-structures-and-algorithms/hash-tables/v/what-is-a-hash-table)

## Abstract Data Types (ADTs)

6.  **Priority Queue:**
    * Implement using a binary heap (min-heap or max-heap).
    * Implement `insert`, `extractMin`/`extractMax`, `peekMin`/`peekMax` operations.
    * **Resources:**
        * [GeeksforGeeks - Priority Queue](https://www.geeksforgeeks.org/priority-queue/)
        * [Khan Academy - Priority Queues](https://www.khanacademy.org/computing/computer-science/algorithms/priority-queues/v/priority-queues)

## Sorting Algorithms

7.  **Basic Sorting Algorithms:**
    * Bubble Sort
    * Selection Sort
    * Insertion Sort
    * **Resources:**
        * [GeeksforGeeks - Sorting Algorithms](https://www.geeksforgeeks.org/sorting-algorithms/)
        * [Khan Academy - Sorting Algorithms](https://www.khanacademy.org/computing/computer-science/algorithms/sorting-algorithms/v/sorting-algorithms)

8.  **Intermediate Sorting Algorithms:**
    * Merge Sort
    * Quick Sort (try implementing different pivot selection strategies)
    * **Resources:**
        * [GeeksforGeeks - Merge Sort](https://www.geeksforgeeks.org/merge-sort/)
        * [GeeksforGeeks - Quick Sort](https://www.geeksforgeeks.org/quick-sort/)
        * [Khan Academy - Merge Sort](https://www.khanacademy.org/computing/computer-science/algorithms/merge-sort/v/merge-sort)
        * [Khan Academy - Quicksort](https://www.khanacademy.org/computing/computer-science/algorithms/quick-sort/v/quick-sort)

9.  **Advanced Sorting Algorithms (Understand the principles):**
    * Heap Sort (you'll have the heap implementation from the Priority Queue)
    * Counting Sort (for integer arrays within a specific range)
    * Radix Sort (for integer arrays)
    * **Resources:**
        * [GeeksforGeeks - Heap Sort](https://www.geeksforgeeks.org/heap-sort/)
        * [GeeksforGeeks - Counting Sort](https://www.geeksforgeeks.org/counting-sort/)
        * [GeeksforGeeks - Radix Sort](https://www.geeksforgeeks.org/radix-sort/)

## Searching Algorithms

10. **Linear Search:** Implement on an array and a linked list.
    * **Resources:**
        * [GeeksforGeeks - Linear Search](https://www.geeksforgeeks.org/linear-search/)
11. **Binary Search:** Implement on a sorted array (iterative and recursive approaches).
    * **Resources:**
        * [GeeksforGeeks - Binary Search](https://www.geeksforgeeks.org/binary-search/)
        * [Khan Academy - Binary Search](https://www.khanacademy.org/computing/computer-science/algorithms/binary-search/v/binary-search)

## Graph Data Structures and Algorithms

12. **Graph Representation:**
    * Adjacency Matrix
    * Adjacency List
    * **Resources:**
        * [GeeksforGeeks - Graph Representation](https://www.geeksforgeeks.org/graph-representation-adjacency-list-and-adjacency-matrix/)
13. **Graph Traversal Algorithms:**
    * Breadth-First Search (BFS)
    * Depth-First Search (DFS) (iterative and recursive approaches)
    * **Resources:**
        * [GeeksforGeeks - Breadth-First Search (BFS)](https://www.geeksforgeeks.org/breadth-first-search-or-bfs-for-a-graph/)
        * [GeeksforGeeks - Depth-First Search (DFS)](https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/)
14. **Basic Graph Algorithms:**
    * Detecting cycles in an undirected graph (using DFS or Union-Find).
    * Detecting cycles in a directed graph (using DFS).
    * **Resources:**
        * [GeeksforGeeks - Detect Cycle in Undirected Graph](https://www.geeksforgeeks.org/detect-cycle-undirected-graph/)
        * [GeeksforGeeks - Detect Cycle in Directed Graph](https://www.geeksforgeeks.org/detect-cycle-in-a-directed-graph-using-colors/)

## Tree Data Structures and Algorithms

15. **Binary Tree:**
    * Implement basic node structure.
    * Implement different tree traversals: Inorder, Preorder, Postorder (recursive and iterative approaches).
    * **Resources:**
        * [GeeksforGeeks - Binary Tree](https://www.geeksforgeeks.org/binary-tree-data-structure/)
        * [GeeksforGeeks - Tree Traversals](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/)
16. **Binary Search Tree (BST):**
    * Implement `insert`, `delete`, `search`, `findMin`, `findMax` operations.
    * **Resources:**
        * [GeeksforGeeks - Binary Search Tree](https://www.geeksforgeeks.org/binary-search-tree-data-structure/)
17. **Balanced Binary Search Trees (Understand the concepts and try a basic implementation):**
    * AVL Trees (focus on understanding rotations)
    * Red-Black Trees (understand the properties and basic insertion/deletion logic)
    * **Resources:**
        * [GeeksforGeeks - AVL Tree](https://www.geeksforgeeks.org/avl-tree-set-1-insertion/)
        * [GeeksforGeeks - Red-Black Tree](https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/)
18. **Heaps (covered in Priority Queue but reiterate here in the context of trees).**
    * **Resources:** (See Priority Queue section)
19. **Tries (Prefix Trees):**
    * Implement `insert`, `search`, `startsWith` operations.
    * **Resources:**
        * [GeeksforGeeks - Trie](https://www.geeksforgeeks.org/trie-insert-and-search/)

## Dynamic Programming (Start with simpler problems)

20. **Fibonacci Sequence:** Implement using recursion with memoization and using iterative bottom-up approach.
    * **Resources:**
        * [GeeksforGeeks - Fibonacci Numbers](https://www.geeksforgeeks.org/program-for-fibonacci-numbers/)
        * [Khan Academy - Dynamic Programming](https://www.khanacademy.org/computing/computer-science/algorithms/dynamic-programming/v/dynamic-programming-introduction)
21. **Coin Change Problem:** Find the minimum number of coins to make a certain amount.
    * **Resources:**
        * [GeeksforGeeks - Coin Change](https://www.geeksforgeeks.org/coin-change/)
22. **Knapsack Problem (0/1 Knapsack):** Given a set of items with weights and values, determine the maximum value that can be put in a knapsack of a given capacity.
    * **Resources:**
        * [GeeksforGeeks - 0/1 Knapsack Problem](https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/)
23. **Longest Common Subsequence (LCS):** Find the longest subsequence common to two sequences.
    * **Resources:**
        * [GeeksforGeeks - Longest Common Subsequence](https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/)

## Advanced Topics (Optional, but good to explore later)

* **Disjoint Set Union (Union-Find):** Implement with path compression and union by rank/size.
    * **Resources:**
        * [GeeksforGeeks - Disjoint Set Data Structures](https://www.geeksforgeeks.org/disjoint-set-data-structures-union-find-algorithm/)
* **Segment Trees:** Understand the structure and implement range query operations.
    * **Resources:**
        * [GeeksforGeeks - Segment Tree](https://www.geeksforgeeks.org/segment-tree-set-1-range-minimum-query/)
* **Fenwick Trees (Binary Indexed Trees):** Understand the structure and implement range sum queries and point updates.
    * **Resources:**
        * [GeeksforGeeks - Fenwick Tree (Binary Indexed Tree)](https://www.geeksforgeeks.org/fenwick-tree-binary-indexed-tree-bit-2/)
* **String Algorithms:**
    * Knuth-Morris-Pratt (KMP) algorithm for pattern searching.
        * [GeeksforGeeks - KMP Algorithm](https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/)
    * Rabin-Karp algorithm for pattern searching.
        * [GeeksforGeeks - Rabin-Karp Algorithm](https://www.geeksforgeeks.org/rabin-karp-algorithm-for-pattern-searching/)
* **Graph Algorithms (more advanced):**
    * Dijkstra's algorithm for finding the shortest path in a weighted graph.
        * [GeeksforGeeks - Dijkstra's Algorithm](https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm/)
    * Bellman-Ford algorithm for finding the shortest path in a weighted graph (can handle negative weights).
        * [GeeksforGeeks - Bellman–Ford Algorithm](https://www.geeksforgeeks.org/bellman-ford-algorithm/)
    * Floyd-Warshall algorithm for finding the shortest paths between all pairs of vertices in a weighted graph.
        * [GeeksforGeeks - Floyd Warshall Algorithm](https://www.geeksforgeeks.org/floyd-warshall-algorithm-dp-28/)
    * Minimum Spanning Tree algorithms (Prim's and Kruskal's).
        * [GeeksforGeeks - Prim's Algorithm](https://www.geeksforgeeks.org/prims-algorithm-greedy-approach/)
        * [GeeksforGeeks - Kruskal's Algorithm](https://www.geeksforgeeks.org/kruskals-minimum-spanning-tree-algorithm-greedy-algorithm-2/)
    * Topological Sort.
        * [GeeksforGeeks - Topological Sorting](https://www.geeksforgeeks.org/topological-sorting/)

Remember to explore these resources and others to deepen your understanding of each topic before and during your implementation. Good luck with your learning journey!