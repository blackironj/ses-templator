# Upload, Download, Delete AWS SES Email template
This is a CLI tool for uploading, downloading, deleting SES Email template easily

Currently, AWS-SES does not provide GUI editor yet. So you can use AWS-CLI tool to manage email-template. ([AWS-CLI-Ref](https://awscli.amazonaws.com/v2/documentation/api/latest/index.html))

OR Use this cli tool

# How to run
```bash
git clone https://github.com/blackironj/ses-templator.git
cd ses-templator
go build .
./ses-templator [command]
```

# Usage
```
ses-templator is a CLI tool.
It helps you update / edit / get templates easily

Usage:
  ses-templator [command]

Available Commands:
  del         delete a SES-email template
  down        download the SES-email template
  help        Help about any command
  list        get list of SES-email templates
  up          upload the SES-email template

Flags:
  -h, --help         help for ses-templator
      --key string   access-key file (default is $HOME/.aws-access-key.json)

Use "ses-templator [command] --help" for more information about a command.
```

- If you do not use some flag, tool use a default value. and some commands can't omit flags

# Get list of template
```
get list of SES-email templates

Usage:
  ses-templator list [flags]

Flags:
  -h, --help      help for list
      --num int   max amount of template (default 10)

Global Flags:
       --key string   access-key file (default is $HOME/.aws-access-key.json)


[run command example]
./ses-templator list --key your/key/path
```

# Get a template
```
download the SES-email template

Usage:
  ses-templator down [flags]

Flags:
  -h, --help                 help for down
  -n, --name string          email template name (required)
  -p, --path template name   download path (default is $HOME/template name.html)

Global Flags:
      --key string   access-key file (default is $HOME/.aws-access-key.json)

[run command example]
./ses-templator down --key your/key/path -n ExampleTemplate -p your/download/file/path
```

# Delete a template
```
delete a SES-email template

Usage:
  ses-templator del [flags]

Flags:
  -h, --help          help for del
  -n, --name string   email template name (required)

Global Flags:
      --key string   access-key file (default is $HOME/.aws-access-key.json)

[run command example]
./ses-templator del --key your/key/path -n ExampleTemplate
```

# Upload a template
```
upload the SES-email template

Usage:
  ses-templator up [flags]

Flags:
  -h, --help          help for up
  -p, --path string   upload template file path (required)

Global Flags:
      --key string   access-key file (default is $HOME/.aws-access-key.json)
    

[run command example]
./ses-templator up --key your/key/path -p your/upload-template-form/path
```

# Config
## aws-access-key.json
```json
{
    "access_key":"your aws-access-key",
    "secret_key":"your aws-secret-key",
    "region": "region ex) us-east-1"
}
```

## upload-template-form.json
```json
{
    "template_name" : "signup_confirm_template",
    "subject" : "Hello {{user}}",
    "html_path": "your/html/template/file/path" 
}
```
