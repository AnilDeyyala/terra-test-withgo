terraform {
  required_providers {
    azurerm = {
        source = "hashicorp/azurerm"
    }
  }
}

provider "azurerm" {
  features {}
}

output "hello_world" {
  value = "Hello, World!"
}

module "gotest_module" {
    source = "../"
    storage_accountname = var.storage_accountname
    location = var.location 
    account_tier = var.account_tier
    rg_name = var.rg_name
    
}


# output "storage_accountname" {
#   value = module.gotest_module.storage_accountname
# }

# output "rg_name" {
#   value = module.gotest_module.rg_name
# }

# output "location" {
#   value = module.gotest_module.location
# }

# output "account_tier" {
#   value = module.gotest_module.account_tier
# }
