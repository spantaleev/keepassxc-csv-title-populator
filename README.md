# KeePassXC CSV export title populator

This utility transforms a CSV file exported by [KeePassXC](https://keepassxc.org/) and populates empty **Title** fields with the parsed hostname from the **URL** field.

If the **URL** field does not contain a valid URL for some reason, `Unknown` is used for the **Title**.

All other columns (besides the **Title** column) are left untouched.


## Reason

When passwords are imported from Mozilla Firefox into KeePassXC, they would have an empty **Title**.

This is generally OK (when using KeePassXC), but may cause trouble when the passwords are further migrated from KeePassXC to another system (like [Bitwarden](https://bitwarden.com/) or [Vaultwarden](https://github.com/dani-garcia/vaultwarden)).

When importing passwords that lack a **Title** into Bitwarden (Vaultwarden), the import would succeed, but records not having a title:

- would be hard to click in the web-vault user interface
- would necessitate that a title (name) is assigned when attempting to edit them

To fix up your KeePassXC database and import it into Bitwarden/Vaultwarden, follow these steps:

1. Export the KeePassXC database to CSV
2. Use this tool to fix up its empty **Title** columns. See [Usage](#usage)
3. Import the newly-created CSV file into a new KeePassXC database. Verify that it looks OK.
4. Export this new KeePassXC database as XML
5. Import the XML file into Bitwarden / Vaultwarden


## Usage

```
go run main.go /path/to/input.csv /path/to/output.csv
```


## Example

### Before

| Group | Title            | Username | Password | URL                            | More columns here |
| ----- | ---------------- | -------- | -------- | ------------------------------ | ----------------- |
| Root  | Some Title       | john     | password | https://example.com            |  ....             |
| Root  |                  | peter    | password | https://example.com            |  ....             |
| Root  |                  | frank    | password | https://another.com            |  ....             |
| Root  |                  | george   | password |                                |  ....             |

### After

| Group | Title            | Username | Password | URL                            | More columns here |
| ----- | ---------------- | -------- | -------- | ------------------------------ | ----------------- |
| Root  | Some Title       | john     | password | https://example.com            |  ....             |
| Root  | example.com      | peter    | password | https://example.com            |  ....             |
| Root  | another.com      | frank    | password | https://another.com            |  ....             |
| Root  | Unknown          | george   | password |                                |  ....             |
