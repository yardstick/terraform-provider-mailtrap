variable api_key {
}

provider mailtrap {
  token = var.api_key
}

resource mailtrap_project test {
  name = "mailtrap_terraform_test"
}

resource mailtrap_inbox test {
  name       = "mailtrap_terraform_test"
  project_id = mailtrap_project.test.id
}

data mailtrap_project default {
  name = "Default"
}

resource mailtrap_inbox default {
  name       = "new_mailbox_in_default"
  project_id = data.mailtrap_project.default.id
}


output host {
  value = mailtrap_inbox.test.smtp_host
}

output username {
  value = mailtrap_inbox.test.smtp_username
}

output password {
  value     = mailtrap_inbox.test.smtp_password
  sensitive = true
}

output port {
  value = 587
}
