#!/bin/bash

output_file="output.csv"
num_lines=100000

echo "id,name,phone,email" > "$output_file"

for ((i=1; i<=num_lines; i++)); do
    random_id=$(uuidgen)
    
    data="name ${i},999999999,email${i}@example.com"
    
    # Concatenar ID com dados e adicionar ao arquivo
    echo "$random_id,$data" >> "$output_file"
done

echo "CSV file created: $output_file"
