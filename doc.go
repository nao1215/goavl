/*
goavl is linter of goa version 1 (URL: https://github.com/shogo82148/goa-v1).
The purpose of development is to speed up Web API design using goa and minimize design
differences within the team. Goa generate the base-files(routing, controller, Swagger, etc.)
required for Web API hosting that based on the design written in DSL.

However, it is difficult to use. If the design fails to generate various files, the error
location cannot be identified immediately (error message is not user friendly). For example,
modern compilers show the number of lines in error. However, goa does not output any such
information.

So, we started developing Linter to instantly identify the problem areas in the goa design
file. goavl will only support goa version 1, will not support the current version (it's
means version 3). The reason is that I'm using forked goa version 1.
*/
package main
