# reverso
simplier reverse proxy for dev environments

- Sample config file:

```json

{
    "handlers": [
        {
            "description": "redirect to ssl",
            "port": 80,
            "type": "proxy",
            "hosts": [
                {
                    "description": "frontend",
                    "host": "local.com",
                    "type": "proxy",
                    "address": "http://local.com:3000"
                },
                {
                    "description": "flask",
                    "host": "api.local.com",
                    "type": "proxy",
                    "address": "http://local.com:5000"
                }
            ]
        }
    ]
}
```

> All description fields are just informative.
