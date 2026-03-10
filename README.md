# chunkd

This is basically my first few steps with go trying to create a backup system for my personal kvm.
Rn ive tried to use a simple chunking logic.

now ive added a simple rebuild logic to rebuild the chunked data, now im planning on creating a function that saves all chunks with hashes as names in a folder and then add a function to compare changes made to the base file to detect modifications


Ive decided to make a python utility or similar and backup my data via borgmatic or restic as for my use case a custom written backup utility seems a bit overkill
