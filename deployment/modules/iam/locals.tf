locals {
  project = var.project
  workload_identity_pool = {
    id          = "${var.env}-github"
    provider_id = "${var.env}-github"
  }
  service_account = {
    github = {
      repository_name = "sugar-cat7/vspo-portal"
    }
  }
}
