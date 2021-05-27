resource "google_container_cluster" "primary" {
    name                = var.cluster_name
    location            = var.region
    remove_default_node_pool    = true
    initial_node_count          = 1
    min_master_version          = var.k8s_version
    
    resource_labels = {
        environment             = "development"
    }
}

resource "google_container_node_pool" "primary_nodes" {
  name               = var.cluster_name
  location           = var.region
  cluster            = google_container_cluster.primary.name
  version            = var.k8s_version
  initial_node_count = "1"
  node_config {
    preemptible  = true
    machine_type = "e2-standard-2"
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
  autoscaling { 
    min_node_count = 1
    max_node_count = 1
  }
  management {
    auto_upgrade = false
  }
  timeouts {
    create = "15m"
    update = "1h"
  }
}

module "iam" {
    source = "./modules/iam"
    
    project_id      = var.project_id

}
