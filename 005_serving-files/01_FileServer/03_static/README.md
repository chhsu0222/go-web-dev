# FileServer special case

**As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".**
Which means, it only serves index.html if it is in the folder.
Otherwise, it serves all the files (including the source codes).


