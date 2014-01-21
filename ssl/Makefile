.PHONY: file.srl
file.srl:
	echo "00" > file.srl


ca.pem: file.srl
	openssl req -keyout ca.key -out ca.pem -new -x509 -config ca.config


server.key:
	openssl genrsa -out server.key 1024


server.req: server.key
	openssl req -key server.key -new -out server.req -config server.config


server.pem: server.req ca.pem
	openssl x509 -req -in server.req -CA ca.pem -CAkey ca.key -CAserial file.srl -out server.pem


client.key:
	openssl genrsa -out client.key 1024


client.req: client.key
	openssl req -key client.key -new -out client.req -config client.config


client.pem: client.req ca.pem
	openssl x509 -req -in client.req -CA ca.pem -CAkey ca.key -CAserial file.srl -out client.pem


serverca.pem: server.pem ca.pem
	cat server.pem > serverca.pem
	cat ca.pem >> serverca.pem


clientca.pem: client.pem ca.pem
	cat client.pem > clientca.pem
	cat ca.pem >> clientca.pem


nick.pem: clientca.pem client.key
	cat clientca.pem > nick.pem
	cat client.key >> nick.pem


.PHONY: all
all: serverca.pem clientca.pem nick.pem


clean:
	rm -rf ca.pem file.srl server.key server.req server.pem client.key client.req client.pem serverca.pem clientca.pem nick.pem ca.key
