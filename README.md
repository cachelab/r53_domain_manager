# Route 53 Domain Manager

CLI tool that will iterate through all of the domains you have registered in
Route 53 and update the Registrant, Administrative and Technical contact
information.

[![CircleCI](https://circleci.com/gh/cachelab/r53_domain_manager.svg?style=svg)](https://circleci.com/gh/cachelab/r53_domain_manager)

## Usage

This cli is configured by the following parameters:

### Examples

```bash
r53dm list

# Parameter is required.
r53dm describe \
--domain=cachelab.co

# All parameters are required and will be updated accordingly.
r53dm update \
--domain=cachelab.co \
--addressline1="1234 Jones Drive" \
--city=Littleton \
--state=CO \
--zip=80127 \
--email=hello@cachelab.co \
--first=Andrew \
--last=Puch \
--phone=+1.1231231234
```

#### Parameters

```bash
domain       Domain you wish to update.
addressline1 Address line 1 you want to change.
city         City you want to change.
state        State you want to change.
zip          Zip you want to change.
email        Email you want to change.
first        First name you want to change.
last         Last name you want to change.
phone        Phone number you want to change.
```
