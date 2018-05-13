# queryserver

## 



## Already cover

### Right of access

- Only a challenge's master can ban/unban players from that challenge and update that challenge.
- Only two players who have a friendship with each other can delete that friendship.
- Players can only change their own post/comment/rating/profile.
- Banned player can not do anything in a challenge and also can not change what they have posted.

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
