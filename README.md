# go-naming-convention
Platform written in GOLang to process input string to specified case of naming conventions based on validity of tokens

Check out the cool terminal to see API in action : https://gocase.azurewebsites.net/

API endpoint : https://gocase-cus.azurewebsites.net/v1/name?token=onetwothree&type=camel
Check this out for more detials : https://github.com/surenderssm/go-naming-convention/blob/master/deploy/content/go-case-namingconventionAPI.v1.docx


![alt text](https://raw.githubusercontent.com/surenderssm/go-naming-convention/master/deploy/content/GoCaseSystemOverview.jpg)

At a high level (and in simplified terms), gocase API convert the text to specified case (format). Naming convention is the core of any programming language or system. Extracting tokens out of given text is the core step in Lexical analysis, first step of compiler/ interpreter.
In our case tokens are extracted based on valid English words.

	token:
o	A free text, made out of only English alphabets
o	At least of length 3, without any space, special character
	type (casetype): 
o	intended naming case which has to be applied after extracting the words
o	Following case are supported in the system

camel	CamelCase https://en.wikipedia.org/wiki/Camel_case

lowercamel	LowerCamelCase https://en.wikipedia.org/wiki/Camel_case  (~CamelCase)

pascal	PascalCase http://wiki.c2.com/?PascalCase  (~UpperCamelCase)

uppercamel	UpperCamelCase https://en.wikipedia.org/wiki/Camel_case (~PascalCase)

snake	SankeCase https://en.m.wikipedia.org/wiki/Snake_case (lowerCase + "_")

darwin	DarwinCase https://en.wikipedia.org/wiki/Camel_case The combination of "TitleCase " and "snake case"

title	TitleCase

lower	LowerCase

upper	UpperCase

