# Biocad-internship

This project was made for an internship at biocad. Main tasks:

* take config (host,port,password, etc.) to connect to the database (MongoDB) as well as the address of the directory
* periodically inspect the directory for new files that have not yet been processed (.tsv) (probably keep a list of those already processed in the database) (see the "Source data" sheet)
* queue file processing 

* put the data from the file in the database
* after processing the file and writing to the database, you need to create a file (rtf,doc,pdf to choose from) with the name from the *unit_guid* field in the input file, with data on this *unit_guid*
* parsing errors (for example, file mismatch) - also write to the database and file
* place output files in a separate directory
* make an API interface that will allow you to get paginated data from the database (page/limit) to get data by *unit_guid*

**Stack: GO, MongoDB, RabbitMQ**
