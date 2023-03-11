# Docker file builds an image with a tinode chat server.
#
# In order to run the image you have to link it to a running database container. For example, to
# to use RethinkDB (named 'rethinkdb') and map the port where the tinode server accepts connections:
#
# $ docker run -p 6060:6060 -d --link rethinkdb \
#	--env UID_ENCRYPTION_KEY=base64+encoded+16+bytes= \
#	--env API_KEY_SALT=base64+encoded+32+bytes \
#	--env AUTH_TOKEN_KEY=base64+encoded+32+bytes \
#	tinode-server

FROM alpine:3.17

