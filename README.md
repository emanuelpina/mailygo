# MailyGo
[![Build Status](https://drone.mnlpn.xyz/api/badges/emanuelpina/mailygo/status.svg)](https://drone.mnlpn.xyz/emanuelpina/mailygo)

MailyGo is a small tool written in Go that allows to send HTML forms, for example from static websites without a dynamic backend, via email. It can be used for contact forms on pages created with [Hugo](https://gohugo.io/) ([example](https://emanuelpina.pt/contact/)).

MailyGo is lean and resource-saving. It can be installed with just one executable file.

## Installation

MailyGo can be compiled with the following command:

```bash
go get -u codeberg.org/emanuelpina/mailygo
```

It can then be executed directly.

## Configuration

To run the server, you must set a few environment variables from the list below.

| Name | Type | Default value | Usage |
|---|---|---|---|
| **`SMTP_USER`** | required | - | The SMTP user |
| **`SMTP_PASS`** | required | - | The SMTP password |
| **`SMTP_HOST`** | required | - | The SMTP host |
| **`SMTP_PORT`** | optional | 587 | The SMTP port |
| **`EMAIL_FROM`** | required | - | The sender mail address |
| **`EMAIL_TO`** | required | - | Default recipient |
| **`ALLOWED_TO`** | required | - | All allowed recipients (separated by `,`) |
| **`PORT`** | optional | `8080` | The port on which the server should listen |
| **`HONEYPOTS`** | optional | `_t_email` | Honeypot form fields (separated by `,`) |
| **`GOOGLE_API_KEY`** | optional | - | Google API Key for the [Google Safe Browsing API](https://developers.google.com/safe-browsing/v4/) |
| **`SPAMLIST`** | optional | `gambling,casino` | List of spam words |
| **`DENYLIST`** | optional | `submit` | List of fields names to deny |
| **`TOKEN`** | optional | - | A token to identify the origin of the submission |
| **`MESSAGE_HEADER`** | optional | - | Text to appear at the beginning of the email message, before the list of fields |
| **`MESSAGE_FOOTER`** | optional | - | Text to appear at the end of the email message, after the list of fields |
| **`MESSAGE_SUBMITTER`** | optional | `false` | If set to `true` and the form submitter provide an email address, a copy of the message is sent to him |
| **`MESSAGE_SUBMITTER_HEADER`** | optional | - | Text to appear at the beginning of the email message sent to submitter, before the list of fields |
| **`MESSAGE_SUBMITTER_FOOTER`** | optional | - | Text to appear at the end of the email message sent to submitter, after the list of fields |

## Special form fields

You can find a sample form in the `form.html` file. Fields whose name do not start with an underscore (`_`) will be sent by email. Fields with an underscore serve as control fields for special purposes:

| Name | Type | Default value | Usage |
|---|---|---|---|
| **`_to`** | optional | - | Recipient, it must be in `ALLOWED_TO`, hidden |
| **`_replyTo`** | optional | - | Email address which should be configured as replyTo, (most probably not hidden) |
| **`_redirectTo`** | optional | - | URL to redirect to, hidden |
| **`_formName`** | optional | - | Name of the form, hidden |
| **`_token`** | optional | - | Token to identify the origin of the submission, hidden |
| **`_t_email`** | optional | - | (Default) "Honeypot" field, not hidden, advised (see notice below) |

As I'm using MailyGo to handle a contact form and I want the fields Name, Subject and Message to be listed on the email in that particular order, I created specific names to handle those fields. **Its use is optional.** You can use all of them or just one. They will be listed on the email just before the others fields.

| Name | Type | Default value | Usage |
|---|---|---|---|
| **`_name`** | optional | - |Form submitter Name |
| **`_subject`** | optional | - |Form submitter Subject |
| **`_message`** | optional | - |Form submitter Message |

## Spam protection

MailyGo offers a set of options to prevent spam:

- A [Honeypot](https://en.wikipedia.org/wiki/Honeypot\_(computing)) field, which is basically another input, but it's hidden to the user with either a CSS rule or some JavaScript. It is very likely, that your public form will get the attention of some bots some day and then the spam starts. But bots try to fill every possible input field and will also fill the honeypot field. MailyGo won't send mails of form submissions where a honeypot field is filled. So you should definitely use it;
- A list of spam words (see **`SPAMLIST`** on [Configuration](#configuration)). If any of the fields include a word from the list the submission will be marked as spam and an email will not be sent;
- A list of fields names to deny (see **`DENYLIST`** on [Configuration](#configuration)). Some bots try send a particular field or set of fields that are not part of the form. This is easy to identify and the name of those fields can de added to this list so MailyGo handle those submissions as spam;
- A token to identify the origin of the submission (see **`TOKEN`** on [Configuration](#configuration) and **`_token`** on [Special form fields](#special-form-fields)). Some bots will *grab* the URL of your MailyGo and post a submission directly to that address, bypassing the form. To prevent this, a token can de used to assure only the submissions that come from the form are handle by MailyGo. This token can be any combination of letters and numbers. If a **`TOKEN`** is defined on configuration, MailyGo will look for a field named **`_token`**. If this field doesn't exist or its value doesn't match the one defined on configuration the submission will be marked as spam and an email will not be sent;
- If a Google Safe Browsing API key is set, submitted URLs will also get checked for threats.

## License

MailyGo is licensed under the MIT license, so you can do basically everything with it, but nevertheless, please contribute your improvements to make MailyGo better for everyone. See the LICENSE file.