# healthy-service


:herb: :mushroom: Little app which allows you to filter vegetables by name and season.

This project is based on Graphql coupled with mongo to handle the data layout.


You can check out the front end here:
https://github.com/pierrebiver/healthy

To launch the project first initialize the db with:
```
mongoimport -d healthy -c foods --drop --jsonArray --file MOCK_DATA-healthy.json
```

Then run main.go file from main folder.

