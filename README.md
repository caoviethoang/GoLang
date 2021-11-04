
# Level 2
### According to the way of the topic gives or want to get the customer id, when reinstalling the application or installing another device, it will be able to display the id of the user that has been ignored again.


### ==> So the solution which I came up with here is: Write a new API to update ignored users, both left / right patterns can use the same API, differentiated by direction (direction / type)

- Add a new table
- Index them
- Bulk insert user_id and user_name of ignored users
- And select "NOT IN" to ensure no duplicate id.
