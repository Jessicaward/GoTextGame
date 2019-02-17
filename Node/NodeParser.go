package Node

/*The whole idea behind this parser is to be able to write a simple file format, which can be parsed by this parser.
The file will need to contain a number of things:
	-Key for the current node.
	-Text to be displayed for the user.
	-Options that the user can select
		-Option text
		-Node 'key' for the option
Example of the file format:
----
:Start <You're in a forest, what do you do?>
>Scream @Scream
>RunAway @Run away

:Scream <You screamed, a bear heard you, and then ate you.>
>Start @Retry?

:RunAway <You ran north, and found a village, everyone is so nice!>
>NextNodeHere @Do something
----

Great Stack Exchange post for this idea: https://gamedev.stackexchange.com/questions/144873/optimizing-data-structure-for-my-text-adventure */