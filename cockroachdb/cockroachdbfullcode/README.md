Structure of this example 

// call by main function 

    cockroachdb 
        |-main.go (called by)
        |-go.sum
        |-go.mod
        |-cockroachdbfullcode
                |- dbConnectionCreate.go (to create database connection and options to select the functionalities)
                |-createTable.go
                |-insertdata.go (include InsertData() function for inserting the data into table)
                |-updatedata.go (include UpdateData() function for updating table in differnt format)
                |-deletedata.go (include DeleteData() function to delete row data from table)
                |-dropcolumn.go (include DropColumn() function to drop the column from table)
                |-droptable.go  (include DropTable() function to drop the table)
                |-truncatetable.go (include TruncateTable() function to truncate table if it exists)
                |-printdata.go (Print table data)