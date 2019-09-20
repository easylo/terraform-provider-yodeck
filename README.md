<br>
<br>


# **[terraform-provider-yodeck](https://github.com/easylo/terraform-provider-yodeck)**

### <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

### Terraform provider for [Yodeck](https://www.yodeck.com/) <span class="colour" style="color:rgb(255, 255, 255)">digital signage </span>

## Using the provider:

##### 00-provider.tf

```
variable "yodeck_username" {
  type    = "string"
  default = "me@domain.gtd"
}

variable "yodeck_apikey" {
  type    = "string"
  default = "MyKEY"
}

provider "yodeck" {
  username = "${var.yodeck_username}"
  apikey   = "${var.yodeck_apikey}"
}

variable "wifi_ssid" {
  type    = "string"
  default = "mywifi"
}

variable "wifi_key" {
  type    = "string"
  default = "connectMe"
}

variable "wifi_mode" {
  type    = "string"
  default = "WEP"
}
```

##### workspace.tf

```
resource "yodeck_workspace" "floor" {
  name        = "Floor workspace"
  description = "Workspace for All Floor"
}
```

##### webpage.tf

```
resource "yodeck_webpage" "home" {
  name        = "home"
  url         = "https://dashboard.domain.gtd/dashboard/home/index.html"
  description = ""
  duration    = 60
  workspace   = "${yodeck_workspace.floor.id}"
}

resource "yodeck_webpage" "admin-home" {
  name        = "home"
  url         = "https://dashboard.domain.gtd/dashboard/admin-home/index.html"
  description = ""
  duration    = 60
  workspace   = "${yodeck_workspace.floor.id}"
}
```

##### playlist.tf

```
resource "yodeck_playlist" "floor" {
  name        = "floor default"
  description = ""
  workspace   = "${yodeck_workspace.floor.id}"

  media {
    media    = "${yodeck_webpage.home.id}"
    duration = "${yodeck_webpage.home.duration}"
  }

  media {
    media    = "${yodeck_webpage.admin-home.id}"
    duration = "${yodeck_webpage.admin-home.duration}"
  }
}
```

##### show.tf

```
resource "yodeck_show" "floor" {
  name      = "Show for floor"
  workspace = "${yodeck_workspace.floor.id}"

  regions {
    top                 = 0
    left                = 0
    height              = 1080
    width               = 1920
    fit                 = "stretch"
    enable_transparency = false
    is_muted            = true
    res_width           = 1920
    res_height          = 1080
    background_audio    = false
    zindex              = 0

    playlists {
      playlist = "${yodeck_playlist.floor.id}"
      duration = 0
      order    = 0
    }
  }
}
```

##### device.tf

```
resource "yodeck_device" "monitor-1" {
    name      = "Monitor 1"
    workspace = "${yodeck_workspace.floor.id}"

    default_show = "${yodeck_show.floor.id}"
    wifi_ssid    = "${var.wifi_ssid}"
    wifi_key     = "${var.wifi_key}"
    wifi_mode    = "${var.wifi_mode}"
}
```

``` sh
terraform plan
```

``` sh
terraform apply
```