# Syntax Rules #

The language rules for the language syntax is defined in a file using a similar format to YACC.
But the definitions are a bit simpler.

The grammar rules defines the parts of the speech that can be parsed by the syntax engine. The
parts of the speech that are used in the dictionary are defined here.

## File Format ##

The file is split in two sections the `token section` and the `syntax rules` section these are
denoted by the `%token` and the `%rules` markers. The token section must precede the rules section.
The rules define tokens, all tokens must be defined before use.

The token and rule names (rule names *are* tokens) must only contain the ascii characters a-z and A-Z,
the numbers 0-9 and can have the '\_'. No whitespace are allowed in the names.

These are valid names:
```
	verb_phrase
	noun
	type44
	Order66	
```

### comments ###

Comments are simply follows the `#` character. Anything the follows this to the end of the line is
discarded. It cannot be escaped.

### token section ###

The leaf tokens (the parts of speech to be decoded) are defined by simply placing the name in the
file either on a single line, or space separated. These must be unique and have the simple name
format.

The token section will follow the `%token` directive.

The following is an example of a valid token section:

```
	%token
	noun, verb, determinate
	pre_position
	adjective, adverb
```	

### syntax rule section ###

The syntax rules are defined in the `%rules` section. The rules are used to connect the tokens to
each other. The rules define the way the sentences are structured and will provide the information
required for the semantic parser to glean the meaning from the utterance.

The syntax rules do not provide any semantic information and quite meaningless statements will be
marked as valid statements, hopefully these will fail when the semantic parsing is done.

The syntax rules allow for optional and repeatable (and both of course) tokens to appear in a clause
definition.

Optional elements are surrounded by the square brackets, i.e. "[adverb]". The repeatable elements are
surrounded by the curly braces, i.e. "{adverb}". The make an element optional and repeatable the curly
braces should be inside the square braces, i.e. "[{adverb}]".

Each clause should be have a unique name, which follows the token format and come at the start of
the line.

The following is a valid syntax rule section:

```
	%rules
	sentence = [noun_phrase] verb_phrase
	noun_phrase = [determinate] [{adjective}] noun
	verb_phrase = [{adverb}] verb [noun_phrase]
```	
