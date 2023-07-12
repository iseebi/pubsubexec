# pubsubexec

Execute specific command when receive Google Cloud Pub/Sub message

## How to use

- Create a Service Account and prepate credential.
    - If run in GCE/Cloud Run, use Application Default Credentials
    - If run in outside of Google Cloud, generate a service account key file and specify the path to it in `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
- Create a Pub/Sub subscription and grant Pub/Sub subscriber privileges to service account
    - Recommend setting the subscription expiration date to indefinite

then start pubsubexec

```
$ export GOOGLE_APPLICATION_CREDENTIALS=path/to/service_account.json
$ pubsubexec \
  --project YOUR_PROJECT_NAME \
  --subscription YOUR_SUBSCRIPTION_NAME \
  --command target_command.sh
```

## License

MIT

