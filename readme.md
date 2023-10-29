# Generating hash


There are two different ways: 
 - md5. For this task sha256 is too much. 
And also we should crop it to 8 chars.
Every time when one FullUrl goes to service its Short part(hash) is the same. 
For simple service is ok, but it's a limit when you decide to add user accounts or some counters(stats) 
 - randInt is my choice, so each FullUrl has a unique ShortUrl pair.