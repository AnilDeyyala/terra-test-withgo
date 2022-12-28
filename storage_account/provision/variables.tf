variable "storage_accountname" {
  type = string
  default = "gotestprovstoarge"
}

variable "account_tier" {
 type = string 
 default = "Standard" 
}

variable "location" {
  type = string 
  default = "East US"
}

variable "rg_name" {
  type = string 
  default = "Devops-RG"
}