# MongoDB Database Schema

## Database `tinode`

### Table `users`
Stores user accounts

Fields:
* `_id` user_id, primary key
* `createdat` timestamp when the user was created
* `updatedat` timestamp when user metadata was updated

Indexes:
 * `_id` primary key
 * `tags` multikey-index (indexed array)
 * `deletedat` index
 * `deviceids` multikey-index of push notification tokens

Sample:
```json

```

### Table `topics`
The table stores topics.

Fields:
* `_id`: name of the topic, primary key
* `createdat`: topic creation time
* `updatedat` timestamp of the last change to topic metadata
* `access` stores topic's default access permission