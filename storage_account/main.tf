resource "azurerm_storage_account" "goteststorage" {
  name                     = var.storage_accountname
  resource_group_name      = var.rg_name
  location                 = var.location 
  account_tier             = var.account_tier
  account_replication_type = "LRS"

  tags = {
    environment = "staging"
  }
}

output "storage_name" {
  value = azurerm_storage_account.goteststorage.name 
}

output "resource_group" {
  value = azurerm_storage_account.goteststorage.resource_group_name
}

output "location" {
  value = azurerm_storage_account.goteststorage.location
}

output "account_tier" {
  value = azurerm_storage_account.goteststorage.access_tier
}