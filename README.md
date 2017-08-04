# multihash
Generate SHA256, SHA1 and MD5 hashes for a set of files, reading the file just once for all hashes. Intended to speed up running sha1sum and md5sum on files one after the other (of course there are bash tricks to do this, but this simplifies it greatly)

## Usage
Simply run the program followed by one or more files:
```bash
./multihash file_one another_file my/folder/third_file
```

To output something similar to this:
```
test1.txt
SHA256: 688787d8ff144c502c7f5cffaafe2cc588d86079f9de88304c26b0cb99ce91c6
SHA1: f10e2821bbbea527ea02200352313bc059445190
MD5: 7815696ecbf1c96e6894b779456d330e

test1 copy.txt
SHA256: 688787d8ff144c502c7f5cffaafe2cc588d86079f9de88304c26b0cb99ce91c6
SHA1: f10e2821bbbea527ea02200352313bc059445190
MD5: 7815696ecbf1c96e6894b779456d330e

test2.txt
SHA256: 17f80754644d33ac685b0842a402229adbb43fc9312f7bdf36ba24237a1f1ffb
SHA1: 153fa238cec90e5a24b85a79109f91ebe68ca481
MD5: 58b4e38f66bcdb546380845d6af27187

```