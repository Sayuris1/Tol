components {
  id: "weaponController"
  component: "/scrpits/weaponController.script"
  properties {
    id: "fire_interval"
    value: "0.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "bullet_speed"
    value: "1000.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/prefabs/bullet.go\"\n"
  ""
}
