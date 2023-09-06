Just for fun. An encoder / decoder for Crockford
base32 messages. Just set the line lenght to 0,
in my original crockford-base32 encoder, to get
properly formatted emoji messages, with a line
length of 35 characters per line.

Usage: cbase32 < message.txt | cee
decode: cee -d < enc_msg.txt | cbase32 -d
