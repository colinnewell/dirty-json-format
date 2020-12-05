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

Is this useful to you?  Who knows.
