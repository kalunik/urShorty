# Generating hash


There are two different ways: 
 - md5. For this task sha256 is too much. 
And also we should crop it to 8 chars.
Every time when one FullUrl goes to service its Short part(hash) is the same. 
For simple service is ok, but it's a limit when you decide to add user accounts or some counters(stats) 
 - randInt is my choice, so each FullUrl has a unique ShortUrl pair.

# Rotation
*WIP

Sometimes in the future it could be critical to store not unique original link with different path as a key.
Maybe query params could help, but for pure link it's kind ugly (imagine you get this from friend 
http://example.com/Hdk201miu?user=cleverfriend1234).

Rotation is our helper. Rotation means that lifetime of link would be limited to certain amount of time.