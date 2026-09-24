# Full-stack demo

The API has an idempotent demo seed mode. It creates the `admin@admin.com` demo
user (password `admin`), a Kubernetes server (`demo-kubernetes` / `demo-secret`),
the `Simple Task Demo` routing (`task`, `sdn_lab_5`) and a stable attempt UUID:
`550e8400-e29b-41d4-a716-446655440000`.

Run it after applying database migrations and before starting Clabgate:

```sh
./backend/apiserver --demo
```

For a container deployment use the same image with command `[/app/apiserver, --demo]`
as a one-shot Job. The regular API then starts without `--demo`.
Configure Clabgate with `CMS_LOGIN=demo-kubernetes`, `CMS_PASSWORD=demo-secret`,
and point `CMS_TASK_URL` at
`https://github.com/maintainer64/cms-labs-simple-task` (branch `main`).

Log in to the frontend with `admin@admin.com` / `admin`, then open the demo
attempt URL. The attempt is intentionally pending until Clabgate acknowledges
the session; this is the same lifecycle used in production.
