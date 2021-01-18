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

