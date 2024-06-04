data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "../../pkg/database/model",
    "--dialect", "mysql",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker+mysql://_/mysql:8.0/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
