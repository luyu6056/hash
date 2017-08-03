# hash
A hash Function for go php js

## GO Example
Newhash().Hash("123","123","32") //81a5f39275fd8f97a79494df034ff6db

Newhash().Hash("123","123","16") //5fvMt5JRhOOb3HyX

Newhash().Hash("123","","32")    //5e82d06f52da6c74464e2f4427d2e733

Newhash().Hash("123","","16")    //W6MnKWks8Gl6dOjp

Newhash().Hash("123")            //W6MnKWks8Gl6dOjp

Newhash().Hash("123","123")      //5fvMt5JRhOOb3HyX

## php Example
echo myhash('123');            //W6MnKWks8Gl6dOjp

echo myhash('123','123');      //5fvMt5JRhOOb3HyX

echo myhash('123123');         //YR60CgOJxbO0B3L0

echo myhash('123','123','32'); //81a5f39275fd8f97a79494df034ff6db

echo myhash('123','','32');    //5e82d06f52da6c74464e2f4427d2e733

## JS Example
The same as php
myhash('123','123','32'); //81a5f39275fd8f97a79494df034ff6db

