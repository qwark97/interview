# Vendor API

### Overview
GET https://api.vendor.com/v1/users

```json
{
  "nextLink": "https://api.vendor.com/v1/users?page=2",
  "users": [
    {
      "id": "1",
      "first_name": "John",
      "last_name": "Doe"
    }
  ]
}
```
### Restrictions
- API is poorly designed and can process only one request at a time
- API returns up to 100 users per request (`size` parameter is not available)
- Request without `page` parameter is treated as `page=1`
- Expected data volume is more than can be processed in memory at once
- When there is no next link, it means all users are fetched
- If request returns HTTP 429, it needs to be retried
- If requests lasts longer than 10 seconds, it needs to be retried
