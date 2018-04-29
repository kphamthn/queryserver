# queryserver

## Need to know

The queryserver works each time I post new documents. But what will happen when we want to update the old 
documents? It has much to do with the right of access, which is not very clear to me right now.

Some examples (only my suggestion):

- Basically an admin can update all documents.
- A normal user can change his profile but not others'.
- A master can edit his own challenge as well as all posts and comments in it. He can also ban players from it.
- A banned player can not post, comment and rate anything in that challenge.
etc.

And there must be some challenges that are hidden from others (that means only certain players can get access to it). 

Did our customers already have a guideline for that, or are they still open?


## Already cover

### Challenge
* Title and description can contain ascii characters.
* Competition mode can only be pvp or ?
* MaxPlayer can only be a number between 1 and 1000.
* Challenge can not end before it starts (end >= start).
* Image can only be an url.
* Master must exist in the database with type "player".

#### Need to know
* How many categories are there?
* How many types of target are there?
* Which value can "completed" have?

### Message
* Message can contain ascii characters.
* Both sender and receiver must exist in the database with type "player".
* Message type can only be "image" or "text"

### Player:
* Email must follow the email format (xxx@yyy.zz).
* Username can contain UTF letters and numbers, have a minumum length of 3 and maximum length of 10.

### Friendship:
* Both player and friend must be in the database with type "player".

### Join:
* Player and challenge must be in the database with respective type.
* "Received" must have a lower value than the end date of the respective challenge.

### Post:
* Description can contain UTF letters and numbers.
* Image must be an url.
* Challenge and Player must exist.

### Rating:
* Rating value must be between -5 and 5.
* All object with "ID" have to exist.

### Comment:
* Post, challenge and player must exist.
* Descripton can contain UTF letters and numbers.
