# Golang Keycloak Example1

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-keycloak-example1.git`
2. Navigate to the folder: `cd go-keycloak-example1`
3. Run the command: `docker compose up`
4. Open new terminal then run command: `go build && ./go-keycloak-example1`
5. Open postman collection file.

#### Test API with no Authorization

```shell
curl --location 'http://127.0.0.1:8081/test' 

401
```

#### Test API with Authorization

```shell
curl --location 'http://127.0.0.1:8081/test' \
--header 'Authorization: Bearer eyJhbGciOiqJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJWNld1ZWcyaUxHb3IxOXF4VFFHWlhmaW5rQTNlNlMwZ2NkSGwzTXc1bGRjIn0.eyJleHAiOjE3MjM2MzMyNjMsImlhdCI6MTcyMzYzMjk2MywianRpIjoiM2RiODIxODEtYWQyMi00YzlkLTk5ZWItNmNmMTJjMGNiODk0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJjM2ZkMmNhYS1hZDk4LTRmMjAtOTYwMS04OWEyZGZkOGEyNzYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxOTIuMTY4LjY1LjEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJzZXJ2aWNlLWFjY291bnQtbXktZ28tc2VydmljZSIsImNsaWVudEFkZHJlc3MiOiIxOTIuMTY4LjY1LjEiLCJjbGllbnRfaWQiOiJteS1nby1zZXJ2aWNlIn0.HprEePbw802bqrsQ24e88ZhSoCiSD2LI_M76FEt0WNgAssVYwq9Da3XMTgn6o192Sm_EmSrt-W9QgckLqFT_hc37tRJsUCsapqCIiiTozqKQoRA8b2Nh0bjXJUoyC-9ei-YJ3lMvTHhRxGVdSjEDHOmQugcipg4YG0WZdGfbm189f6S0q3mSIDbchdXodB0XWUbdITGiTTfb-0NOw-R0RyaRIxaSKueJ6d7i-cJWwEKBIfExUDMsLswStruaGkSJpKLMp-iVsVJDUPl5gqTjXwzbtgi-4syy6iGd5C-iIfnXRhEPeWwAWcrcXhQ6fbhoJmf5-ncLyGWxCCKJpT6QOg'

{
  OK
}
```

#### Test API with Authorization

```shell
curl --location 'http://127.0.0.1:8081/docs' \
--header 'Authorization: Bearer eyJhbGciOiqJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJWNld1ZWcyaUxHb3IxOXF4VFFHWlhmaW5rQTNlNlMwZ2NkSGwzTXc1bGRjIn0.eyJleHAiOjE3MjM2MzMyNjMsImlhdCI6MTcyMzYzMjk2MywianRpIjoiM2RiODIxODEtYWQyMi00YzlkLTk5ZWItNmNmMTJjMGNiODk0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJjM2ZkMmNhYS1hZDk4LTRmMjAtOTYwMS04OWEyZGZkOGEyNzYiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6ImVtYWlsIHByb2ZpbGUiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxOTIuMTY4LjY1LjEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJzZXJ2aWNlLWFjY291bnQtbXktZ28tc2VydmljZSIsImNsaWVudEFkZHJlc3MiOiIxOTIuMTY4LjY1LjEiLCJjbGllbnRfaWQiOiJteS1nby1zZXJ2aWNlIn0.HprEePbw802bqrsQ24e88ZhSoCiSD2LI_M76FEt0WNgAssVYwq9Da3XMTgn6o192Sm_EmSrt-W9QgckLqFT_hc37tRJsUCsapqCIiiTozqKQoRA8b2Nh0bjXJUoyC-9ei-YJ3lMvTHhRxGVdSjEDHOmQugcipg4YG0WZdGfbm189f6S0q3mSIDbchdXodB0XWUbdITGiTTfb-0NOw-R0RyaRIxaSKueJ6d7i-cJWwEKBIfExUDMsLswStruaGkSJpKLMp-iVsVJDUPl5gqTjXwzbtgi-4syy6iGd5C-iIfnXRhEPeWwAWcrcXhQ6fbhoJmf5-ncLyGWxCCKJpT6QOg'

[
    {
        "id": "1",
        "num": "ABC-123",
        "date": "2024-08-14T10:57:15.579606Z"
    },
    {
        "id": "2",
        "num": "ABC-456",
        "date": "2024-08-14T10:57:15.579606Z"
    }
]
```