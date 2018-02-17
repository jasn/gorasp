# gorasp
Rank Select structure on bitarrays in golang.

## Rank and Select data structure problem
In the rank/select problem we are to store a bit array of length n while supporting the operations rank and select.
Additionally we also need to be able to retrieve the i'th bit.
After defining what rank and select is, it will become apparant that this is a static data structuring problem, i.e. the bits do not change.
There is also research on maintaining a bit array dynamically and supporting these operations, but we do not deal with that here.

* `RankOfIndex(i)`: Called Rank in the literature. return the number of 1-bits to the left of the i'th bit.
* `IndexWithRank(i)`: Called Select in the literature. return the lowest index j where RankOfIndex(j) = i.

In the research literature there are solutions that store n bits + o(n) bits, i.e. a lower order term, while supporting the operations in constant time.

## Motivation
Using bit-arrays we can represent many "things" in small space.
When we can get away with using less space, we can have larger structures in memory, thus increasing our throughput.
The most common example is probably representing binary trees using 2n bits, instead of the usual nlog n bits of pointer structures.
But to make it useful we also need to support navigation.
It turns out that this is indeed possible with O(1) rank/select calls for trees.

## Our solution
In our solution we use 1.5n bits in addition to the n bits.
We use n bits, where for each 64 bits, we store a number that is a count of the number of 1-bits up to that position.
We use 0.5n bit where for each i that is a multiple of 64, we store the answer to select(i).
Finally we store the n bits.

When answering a rank query, we only need to find the number of 1-bits in a word after looking up the prefix count.
To answer a select query, we look up the nearest answer, and start scanning words until we find the right index.

Note that this means the select query is not guaranteed to run in O(1) time, as there could be very far between the stored answer.
I.e. select(j+1) - select(j) could be large.
However if the bitarrays we encounter are somewhat balanced, then this is likely not ever an issue.
One obvious improvement is to perform a binary search instead of a linear scan, which could help resolve this issue, and give a guarantee of O(log n) query time.


### Improvements for the future
Currently the space usage is quite high.
We could definitely engineer this to be much much smaller.
For instance, instead of storing an answer for every 64 bits, do it for every 512 bits.
We could also just add more levels, and only store a prefix sum within a level.
In that way we can reduce the number of bits needed to store an answer, and only use 64 bits on the top level, and have sufficiently far spread out.

I already mentioned we might change select to perform a binary search.
But we could also utilize that when we do a lookup in an array we receive 512bits.
That suggest we should store a bit more information that can guide our search better.

It would also be nice, if we could change the function 'selectInWord' in 'rank_select_fast.go' to perform a binary search in the word, and when only 8 bits remain perform a lookup in a precomputed table.
