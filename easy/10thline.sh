# Read from the file file.txt and output the tenth line to stdout.
#!/bin/bash
input="10thline.txt"
LNU=0
while IFS= read -r var
do
  let LNU=$LNU+1
  if [ $LNU -eq 10 ]
  then
      echo "$var"
  fi
done < "$input"
