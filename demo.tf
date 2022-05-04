provider "meshapi" {
    url = "https://DEMO.meshstack.io"
    headers = {
        Authorization = "BASE64_VAL"
    }
}

resource "meshapi_mesh_customer" "demo_customer" {
    name = "demo_customer"
    display_name = "My Demo Customer"
    tags = jsonencode(
        {
            CustomerContact = "AdminUser"
        }
    )
}

resource "meshapi_mesh_project" "demo_app" {
    name = "demo_app"
    display_name = "My Demo App"
    customer_id = meshapi_mesh_customer.demo_customer.id
    tags = jsonencode(
        {
            ProjectContact = "DeveloperUser"
            CostCenter = 0000
        }
    )
}

// GET DEMO USER'S DATA TO SHOW DATA USAGE
data "meshapi_mesh_user" "demo" {
    name = "DEMO@meshcloud.io"
}

resource "meshapi_mesh_customer_user_binding" "demo_user_customer_owner_access" {
    role_name = "Customer Owner"
    customer_id = meshapi_mesh_customer.demo_customer.id
    user_id = data.meshapi_mesh_user.demo.id
}

resource "meshapi_mesh_project_user_binding" "demo_user_project_owner_access" {
    role_name = "Project Owner"
    customer_id = meshapi_mesh_customer.demo_customer.id
    project_id = meshapi_mesh_project.demo_app.id
    user_id = data.meshapi_mesh_user.demo.id
}