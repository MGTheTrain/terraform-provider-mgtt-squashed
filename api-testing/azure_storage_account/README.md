# Azure storage account handler

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Cli applcation for managing Azure storage account services and Azure storage account service containers.  

## References

- [Create an Azure Storage account with the REST API](https://learn.microsoft.com/en-us/rest/api/storagerp/storage-sample-create-account)
- [Storage Accounts API documentation](https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts)

## How to use

### Preconditions

You need an Azure subscription, and you must have Owner permissions for that subscription. Optionally you can utilize a service principal. You also need to [install the az cli tool](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli).

### Authenticate with az cli

If multi-factor authentication is enabled run following:

```sh
az login --tenant <your tenant id>
```

A browser window will appear, prompting you to log in with the Microsoft authentication provider. 

Optionally utilize a service principal and authenticate in a terminal with:

```sh
az login --service-principal -u <your service principal client id> -p <your service principal client secret> --tenant <your tenant id>
```

### Display your bearer token after successfull authentication

```sh 
az account get-access-token 
```

### Build the cli application

On Windows OS:

```sh
 go build -o azure_storage_account_handler.exe
```

On Unix systems (MacOS or Ubuntu >=18.04/debian >=11):

```sh
 go build -o azure_storage_account_handler
```

### Run the cli application

```sh
# [C]reate a storage account service
# Adjust the `assets/create_azure_storage_account_request_body.json` file
azure_storage_account_handler storage_account create -s <subscription_id> -g <resource_group_name> -a <storage_account_name> -t <bearer_access_token> -r <req_body_json_file>

# [R]ead a storage account service by storage account name
azure_storage_account_handler storage_account get -s <subscription_id> -g <resource_group_name> -a <storage_account_name> -t <bearer_access_token>

# [U]pdate a storage account service by storage account name
azure_storage_account_handler storage_account update -s <subscription_id> -g <resource_group_name> -a <storage_account_name> -t <bearer_access_token> -r <req_body_json_file>

# [D]elete storage account service by storage account name
azure_storage_account_handler storage_account delete -s <subscription_id> -g <resource_group_name> -a <storage_account_name> -t <bearer_access_token> 
```