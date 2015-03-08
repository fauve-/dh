This is a really simple pair of programs I worked out on a saturday afternoon.

It's an echo server and client that perform the diffie hellman key exchane, and then use the resulting set of bytes as the initialization of an rc4 stream cypher.
Essentially, the client encodes the input, the server receives it, then decodes and reencodes it, and sends it back to the client, who decodes the payload and prints it to the shell. The client program's input is deliniated by the newline '\n'.

Another interesting thing is that the entire program was written as a single module, but the functionality is decided at runtime by a variable set during link time. A fun trick, but the use of packages would probably be better. But saturday afternoon projects are all about the fun.