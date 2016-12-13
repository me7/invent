<!-- vscode-markdown-toc -->
	* 1. [feature](#feature)

<!-- /vscode-markdown-toc --># invent
Inventory Manager
* experiment project that want to familiar with go standard library
* net/http
* html/template
* JSON
* some db (MongoDb or bolt)
* send email for notification
* some front-end stuff (D3)

###  1. <a name='feature'></a>feature
* inventory create/update/count
* notification when qty low / schedule maintenance
* inventory group - notify when ready
* REST API (add/update/list item)
* user/role management


## references
* [PHP-Laravel](github.com\stevebauman\inventory\)
  * add metric (unit qty), category, location then add item which link to location and category
  * have stock_movement record
  * each item can have several supplier
  * user have to login
* [Golang](github.com\fritz0705\inventory\)
  * user have to login
  * have part which link to category, place and owner
  * have separate part_amount to show amount and each timestamp
  * each item can have attachment
