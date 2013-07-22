Graph
======

Package `graph` is an implementation of various graphs.

Status
======

Everything in [here is tested](coverage.md).  However there are still
essential parts missing, such as a symbolic implementation of each graphs
and APIs to load the graphs with data from an `io.Reader`. This will be
implemented in a not too far future.

This library has also not been optimized.  However, the graphs can handle
sizes in the hundred millions vertices.  Some algorithms will be very slow
on such sizes, most will have an acceptable running time.

Credits
=======

Pretty much all of the algorithms are taken from Algorithm 4th Ed, by
Sedgewick and Wayne.  The book is written in Java, I've ported it to Go
as an exercise in studying the algorithms.

I hope to one day add more algorithms that could make use of the concurrent
capabilities of Go.

License
=======

The [MIT License](LICENSE).
