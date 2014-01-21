skype-irc-bouncer
=================

Proxy and translate Skype-IRC communication. That is, use an IRC client to 
chat with your Skype contacts.





Setup
=====

  - populate the ./ssl directory
   - (from: http://acs.lbl.gov/~boverhof/openssl_certs.html)
   - echo "00" >> file.srl
   - openssl req -out ca.pem -new -x509 
   - openssl genrsa -out server.key 1024 
   - openssl req -key server.key -new -out server.req 
   - openssl x509 -req -in server.req -CA ca.pem -CAkey privkey.pem -CAserial file.srl -out server.pem 
   - openssl genrsa -out client.key 1024 
   - openssl req -key client.key -new -out client.req 
   - openssl x509 -req -in client.req -CA ca.pem -CAkey privkey.pem -CAserial file.srl -out client.pem
   - cat server.pem > serverca.pem
   - cat ca.pem >> serverca.pem
   - cat client.pem > clientca.pem
   - cat ca.pem >> clientca.pem
   - cat clientca.pem > nick.pem
   - cat client.key >> nick.pem
  - configure IRC client certs http://www.oftc.net/NickServ/CertFP/
