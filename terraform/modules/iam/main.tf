terraform {
    required_version = ">= 0.12"
}

resource "google_service_account" "prod-ci" {
    account_id      = "prod-ci"
    display_name    = "prod-ci"
    description     = "Service account for production"
    project         = var.project_id
}

resource "google_project_iam_custom_role" "prod-ci-role" {
    project         = var.project_id
    role_id         = "prodCiRole"
    title           = "Role for production"
    description     = "Role for production usage"
    permissions     = ["iam.roles.list"]
}

resource "google_project_iam_member" "prod-ci-role_member" {
    project         = var.project_id
    role            = "projects/${var.project_id}/roles/${google_project_iam_custom_role.prod-ci-role.role_id}"
    member          = "serviceAccount:${google_service_account.prod-ci.email}"
}
