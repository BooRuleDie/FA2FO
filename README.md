# Running Stripe CLI

You can use this command to run the stripe cli, after it's run successfully you can get the endpoint secret as well. which is a required secret for payment webhook to function properly.
```bash
stripe listen --forward-to localhost:8082/webhook

# Ready! Your webhook signing secret is 'whsec_<REDACTED>' (^C to quit)
```