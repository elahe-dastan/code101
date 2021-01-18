# Securely transmitting data (beginner guide)
I assume you are a person knowing nothing about how to transmit data securely, so I'm going to explain all you need step 
by step.

## What the heck do we mean by secure?
As you know there are always loopholes which the bad guys can use to sniff your confidential information like your bank 
account information and so on. what we want to do is like talking in a language which no one understands but us, this 
way we can talk to each other loudly in the crowd with no anxiety. Wait, what if another person learns our language too?

## Symmetric vs Asymmetric encryption
Symmetric cryptography is an old approach in which there is only one key. The sender encrypts the data using the key, and 
the receiver uses the same key to decrypt the data. The algorithm itself works fine but there is a huge bug, how do you 
want to give the key to the other side who you want to transmit information to? That's where asymmetric cryptography walks
in. In this method you have two keys a private one which you give to nobody, and a public one which you can put on your
instagram bio. A message encrypted by a private key can be decrypted by its public key, and a message encrypted by a 
public key can be decrypted by its private key so if a person picks up your public key from your bio, writes a message and
encrypt it using the key no one can decrypt the message but you

## Digital Certificates
Now let's talk about a real situation suppose you see a pair of sneakers on Amazon and want to order them, to pay for them
you give all the information of your credit card and not worry at all, clearly there is an asymmetric cryptography, you 
encrypt your information by a public key and only the Amazon site has the private key and is able to decrypt but when did
you get the public key?!! That's why we need digital certificates like SSL and TLS. 

## TLS/SSL
Note that TLS is just an updated version of SSL. SSL operates directly on top of the TCP and allows higher protocol layers
to remain unchanged. After using an SSL certificate, all an attacker will be able to see is IP, port and length of data. 
They may be able to terminate the connection but both sides understand this has been done by a third party. Let's go through
what happens step by step:<br/>
1. SSL starts to work after the TCP connection is established, initiating what is called an SSL handshake.<br/>
2. The server sends its certificate to the user along with a number of specifications(including which version of SSL/TLS
   and which encryption methods to use,etc.).
3. The user then checks the validity of the certificate, and selects the highest level of encryption that can be supported 
by both parties and starts a secure session using these methods.There are a good number of sets of methods available with 
   various strengths they are called cipher suites.
4. To guarantee the integrity and authenticity of all messages transferred, SSL and TLS protocols also include an 
   authentication process using message authentication codes(MAC).
   
## RSA
Up to here, I have not said a word about how the encryption, and the decryption is done. Well, there are many algorithms 
but one of the most famous ones is RSA. As you can guess, RSA is asymmetric. You can easily find the algorithm 
implementation on Google here I only mention the idea. The idea of RSA is based on the fact that it is difficult to factorize 
a large integer. The public key consists of two numbers where one number is multiplication of two large prime numbers and 
private key is also derived from the same two prime numbers. So if somebody can factorize the large number, the private 
key is compromised. Therefore, encryption strength totally lies on the key size and if we double or triple the key size, 
the strength of encryption increases exponentially. RSA keys can be typically 1024 or 2048 bits long.

## SHA-256
SHA-256 is a hash function that transforms any text to a 32-byte string. A hash is not encryption because it cannot be 
decrypted back to the original text. It is a one-way cryptographic function, and is a fixed size for any size of source 
text. Hashing is used in cases like `challenge handshake authentication`, `anti-tamper` and `digital signatures`. Note
that hash functions are not appropriate for storing encrypted passwords, as they are designed to be fast to compute, and 
hence would be candidates for brute-force attacks.

## JWT
### Authorization vs Authentication
In authentication, you get a username and password and authenticate that they are correct. It's like logging a user in 
but in authorization you want to make sure that the user that sends request to your server is the same user who logged in
during authentication.
### JWT vs SessionId
Long time ago, people would use session ids for authorization. After logging in, the server would generate a unique id
for the user, keeps the id in the memory and give it to the user, the user keeps the id in its cookies and sends it back
to the server whenever it has another request, the server then checks for the id in its memory. This approach has two 
significant disadvantages, first, a lot of memory should be used in the server side to store the session ids, second,
if you want to have several servers to increase high availability of your application, the servers should share their 
memory which is not always the case. Nowadays, when a user logs in, the server creates a json web token(JWT), it actually
encodes information and signs it with its own secret key then sends the token to the client, the client will send back the 
token to the server in case of any new request, and the server can verify the JWT token, so you see, nothing needs to be 
stored in the server side.
### How JWT Works
