# preset.json example to be used

`{"token":"user_token", "author":"user_id", "channel":["target_id_1", "target_id_2"]}`

## how to get author/user_id
1. Open discord in browser
2. Enable developer mode
3. Right click on yourself -> Copy User ID


## how to get channel/target
1. Open discord in browser
2. Enable developer mode
3. Right click on DM/Server -> Copy Channel ID / Copy Server ID

## how to get token
1. Open discord in browser
2. Open Web Developer Tools
3. Select network section
4. Reload discord page to get list of requests
5. Select a request, for example https://discord.com/api/v9/science
6. From Headers section copy Authorization token
