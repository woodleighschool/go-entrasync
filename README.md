# go-entrasync

A thin Microsoft Graph reader for Entra.

Fetches users, security groups, and transitive group memberships via the Graph API.
Persistence is the caller's responsibility.

## Usage

```go
import entrasync "github.com/woodleighschool/go-entrasync"

client, err := entrasync.NewClient(entrasync.Config{
    TenantID:     "...",
    ClientID:     "...",
    ClientSecret: "...",
})
if err != nil {
    return err
}

snapshot, err := client.Snapshot(ctx)
// snapshot.Users, snapshot.Groups, snapshot.Members
```

## Config

| Field                    | Default  | Description                                     |
| ------------------------ | -------- | ----------------------------------------------- |
| `TenantID`               | required | Azure tenant ID                                 |
| `ClientID`               | required | App registration client ID                      |
| `ClientSecret`           | required | App registration client secret                  |
| `GroupMemberConcurrency` | 8        | Parallel group member fetches during `Snapshot` |

## Versioning

Tagged releases follow semver. Use `go get github.com/woodleighschool/go-entrasync@latest` for the latest release.
