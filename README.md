
**annoy-go** is a version of [annoy](https://github.com/spotify/annoy/) that was generated for Golang according [this instruction](https://github.com/spotify/annoy/blob/master/README_GO.rst).

Warning: it has memory leaks.

To reproduce memory leaks scenario:
 * install docker and docker-compose
 * run `docker-compose run stage`
 * watch how it slowly eats all your RAW

