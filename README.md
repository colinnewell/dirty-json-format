# dirty-json-format

Build/install:

    make && sudo make install

Partial json formatter using the go standard lib 'Indent' function.

    cat test.json | djf

Just a crude CLI wrapped around it to allow you to use it in pipes.  The unique
selling point is that it's a hacked version of the standard library function
which allows the json to be incomplete or malformed.

Note that it will only emit up as much of the json as it understands.

For example:

    echo '{"a":3,mkdsamkdsa"dmsksdm"}' | djf

Will only emit as much of that as is valid:

    {
      "a": 3,

Is this useful to you?  Who knows.  My primary use case is for truncated
json, so I can at least read what I have.  I guess if json is garbled it is
likely to highlight exactly where as it will stop rendering when it gets to
the broken bit.

Why on earth did I copy the source from the standard library?  I couldn't
spot a simple way to extend it.  It's all beautifully formed, but my Go fu
isn't strong enough to know how to alter the function to prevent the
truncation of the buffer passed to Indent to allow me to display what could
be processed.  In theory I ought to contribute something back to the
original library to allow this, but I'm not sure how you would logically
extend this.  It would be neater if an interface was passed in rather than
the Buffer struct, or something along those lines.

There's so much useful code, but it's not exported, and for niche features
like this I'm not sure if there is a clear and obvious way people would be
happy with the core library to be expanded.  I can understand why you don't
export, every visible surface is something you have to continue supporting,
but there's just so much just under the surface that would be so useful :)

Anyway, this is a neat way to prototype the idea without actually messing
with the core lib.
