It's always hard to remember how to curl.<br/>
When content type is json I use the README of url shortener repository, but when content type is form use the below 
curls examples:<br/>

curl -v -F "kind=profile" -F "file=@test.png" -H 'Header-Data: {"cellphone": "1373"}' 127.0.0.1:8080/api/



For future test only<br/>
UPDATE applicants SET Documents = '{"sf":"dfchf"}' WHERE cellphone="1373";